package impl

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	log "github.com/sirupsen/logrus"
	"time"
)

type authorizationService struct {
}

func (a *authorizationService) CheckRecordAccess(ctx context.Context, params service.CheckRecordAccessParams) errors.ServiceError {
	if util.IsSystemContext(ctx) {
		return nil
	}

	userDetails := util.GetUserDetailsFromContext(ctx)

	if userDetails == nil && !annotations.IsEnabled(params.Resource, annotations.AllowPublicAccess) {
		return errors.AccessDeniedError.WithDetails("Public access is denied")
	}

	//var constraints []resource_model.SecurityConstraint
	//
	//for _, constraint := range params.Resource.SecurityConstraints {
	//	constraint.Resource = params.Resource.Name
	//	constraint.Namespace = params.Resource.Namespace
	//}

	//constraints = append(constraints, params.Resource.SecurityConstraints...)
	//constraints = append(constraints, userDetails.SecurityConstraints...)

	//errorFields := a.evaluateConstraints(ctx, params, constraints, userDetails)
	var errorFields []*model.ErrorField

	if len(errorFields) == 0 {
		return nil
	} else {
		return errors.AccessDeniedError.WithDetails("User don't have permission to access this resource").WithErrorFields(errorFields)
	}
}

func (a *authorizationService) evaluateConstraints(ctx context.Context, params service.CheckRecordAccessParams, constraints []resource_model.SecurityConstraint, userDetails *abs.UserDetails) []*model.ErrorField {
	logger := log.WithFields(logging.CtxFields(ctx))

	now := time.Now()
	var errorFields []*model.ErrorField

	/*
	  Default policy for checking constraints are like that
	  1. If anyone rejects, then reject
	  2. If none rejects and anyone allows, then allow
	*/

	// Default permit type is disallow
	hasAllowFlag := false

	var remainingPropertyCheck = make(map[string]bool)
	for _, property := range params.Resource.Properties {
		remainingPropertyCheck[property.Name] = true
	}

	for _, constraint := range constraints {
		logger.Tracef("Evaluating constraint: %v", constraint)
		hasAllowFlagLocal, errorFieldsLocal := a.evaluateConstraint(ctx, params, constraint, now, userDetails, &remainingPropertyCheck)

		logger.Tracef("Constraint evaluation result: %v, %v", hasAllowFlagLocal, errorFieldsLocal)

		if hasAllowFlagLocal {
			hasAllowFlag = true
		}

		if errorFieldsLocal != nil {
			errorFields = append(errorFields, errorFieldsLocal...)
		}
	}

	// check remaining properties
	for property, matched := range remainingPropertyCheck {
		if !matched {
			errorFields = append(errorFields, &model.ErrorField{
				Property: property,
				Message:  fmt.Sprintf("Property '%s' is not allowed", property),
			})
		}
	}

	// if none rejects and anyone allows, then allow

	if !hasAllowFlag {
		errorFields = append(errorFields, &model.ErrorField{
			Property: "resource",
			Message:  "No constraints matched",
		})
	}

	return errorFields
}

func (a *authorizationService) evaluateConstraint(ctx context.Context, params service.CheckRecordAccessParams, constraint resource_model.SecurityConstraint, now time.Time, userDetails *abs.UserDetails, remainingPropertyCheck *map[string]bool) (bool, []*model.ErrorField) {
	logger := log.WithFields(logging.CtxFields(ctx))

	// check resource constraint matches

	if constraint.GetResource() != "*" && constraint.GetResource() != params.Resource.Name {
		// skipping as not related to this resource
		logger.Tracef("Skipping constraint as not related to this resource: %v", constraint)
		return false, nil
	}

	if constraint.GetNamespace() != "*" && constraint.GetNamespace() != params.Resource.Namespace {
		// skipping as not related to this namespace
		logger.Tracef("Skipping constraint as not related to this namespace: %v", constraint)
		return false, nil
	}

	if constraint.GetRecordIds() != nil {
		var found = false

	mainLoop:
		for _, id := range constraint.GetRecordIds() {
			id = a.processValue(id, userDetails)

			if params.Records == nil {
				logger.Tracef("Skipping constraint as records not found: %v", constraint)
				return false, nil
			}

			for _, record := range *params.Records {
				if record.Id == id {
					found = true
					break mainLoop
				}
			}
		}

		if !found {
			logger.Tracef("Skipping constraint as record id not matched: %v", constraint)
			return false, nil
		}
	}

	if constraint.GetOperation() != resource_model.SecurityConstraintOperation_FULL && string(constraint.Operation) != params.Operation.String() {
		logger.Tracef("Skipping constraint as operation not matched: %v", constraint)
		return false, nil
	}

	if constraint.Before != nil && constraint.Before.After(now) {
		logger.Tracef("Skipping constraint as before time not matched: %v", constraint)
		return false, nil
	}

	if constraint.After != nil && constraint.Before.After(now) {
		logger.Tracef("Skipping constraint as after time not matched: %v", constraint)
		return false, nil
	}

	if constraint.Username != nil && *constraint.Username != "*" && *constraint.Username != userDetails.Username {
		logger.Tracef("Skipping constraint as username not matched: %v", constraint)
		return false, nil
	}

	if constraint.Role != nil && *constraint.Role != "*" && !util.ArrayContains(userDetails.Roles, *constraint.Role) {
		logger.Tracef("Skipping constraint as role not matched: %v", constraint)
		return false, nil
	}

	if constraint.Property != "*" {
		if constraint.PropertyMode != nil && *constraint.PropertyMode == resource_model.SecurityConstraintPropertyMode_PROPERTYMATCHONLY {
			if constraint.PropertyValue != nil {
				for _, record := range *params.Records {
					for key := range record.Properties {
						if key == "id" {
							continue
						}
						if key != constraint.Property {
							continue
						}

						(*remainingPropertyCheck)[key] = false

						if constraint.PropertyValue != nil {
							var value = *constraint.PropertyValue

							value = a.processValue(value, userDetails)

							strActualVal := fmt.Sprintf("%v", record.Properties[constraint.Property].AsInterface())

							if strActualVal != value {
								logger.Tracef("Skipping constraint as property value not matched: %v", constraint)
								return false, nil
							}
						}
					}
				}
			}
		} else {
			var value = *constraint.PropertyValue

			value = a.processValue(value, userDetails)

			if params.Records != nil {
				for _, record := range *params.Records {
					if record.Properties[constraint.Property] == nil {
						logger.Tracef("Skipping constraint as property not found: %v", constraint)
						return false, nil
					}

					strActualVal := fmt.Sprintf("%v", record.Properties[constraint.Property].AsInterface())

					if strActualVal != value {
						return false, []*model.ErrorField{
							{
								Property: constraint.Property,
								Message:  fmt.Sprintf("Property '%s' is not allowed", constraint.Property),
							},
						}
					}
				}
			}
		}
	}

	return true, nil
}

func (a *authorizationService) processValue(value string, userDetails *abs.UserDetails) string {
	var processedValue = value

	if processedValue == "$userId" {
		processedValue = userDetails.UserId
	} else if processedValue == "$username" {
		processedValue = userDetails.Username
	}

	return processedValue
}

func NewAuthorizationService() service.AuthorizationService {
	return &authorizationService{}
}
