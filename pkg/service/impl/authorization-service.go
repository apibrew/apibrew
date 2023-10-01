package impl

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/resource_model"
	"github.com/apibrew/apibrew/pkg/resources"
	"github.com/apibrew/apibrew/pkg/service"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/apibrew/apibrew/pkg/util/jwt-model"
	log "github.com/sirupsen/logrus"
	"time"
)

type authorizationService struct {
}

func (a *authorizationService) CheckIsExtensionController(ctx context.Context) errors.ServiceError {
	if err := a.CheckRecordAccess(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ExtensionResource,
		Operation: resource_model.PermissionOperation_FULL,
	}); err != nil {
		return err
	}

	return nil
}

func (a *authorizationService) CheckRecordAccess(ctx context.Context, params service.CheckRecordAccessParams) errors.ServiceError {
	if util.IsSystemContext(ctx) {
		return nil
	}

	userDetails := jwt_model.GetUserDetailsFromContext(ctx)

	if userDetails == nil {
		if !annotations.IsEnabled(params.Resource, annotations.AllowPublicAccess) {
			return errors.AccessDeniedError.WithDetails("Public access is denied")
		} else {
			return nil
		}
	}

	var permissions []*resource_model.Permission

	//for _, permission := range params.Resource.Permissions {
	//	permission.Resource = params.Resource.Name
	//	permission.Namespace = params.Resource.Namespace
	//}

	//permissions = append(permissions, params.Resource.Permissions...)
	permissions = append(permissions, userDetails.Permissions...)

	errorFields := a.evaluateConstraints(ctx, params, permissions, userDetails)

	if len(errorFields) == 0 {
		return nil
	} else {
		return errors.AccessDeniedError.WithDetails("User don't have permission to access this resource").WithErrorFields(errorFields)
	}
}

func (a *authorizationService) evaluateConstraints(ctx context.Context, params service.CheckRecordAccessParams, permissions []*resource_model.Permission, userDetails *jwt_model.UserDetails) []*model.ErrorField {
	logger := log.WithFields(logging.CtxFields(ctx))

	now := time.Now()
	var errorFields []*model.ErrorField

	/*
	  Default policy for checking permissions are like that
	  1. If anyone rejects, then reject
	  2. If none rejects and anyone allows, then allow
	*/

	// Default permit type is disallow
	hasAllowFlag := false

	var remainingPropertyCheck = make(map[string]bool)
	for _, property := range params.Resource.Properties {
		remainingPropertyCheck[property.Name] = true
	}

	for _, permission := range permissions {
		logger.Tracef("Evaluating permission: %v", permission)
		hasAllowFlagLocal, errorFieldsLocal := a.evaluateConstraint(ctx, params, permission, now, userDetails, &remainingPropertyCheck)

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
			Message:  "No permissions matched",
		})
	}

	return errorFields
}

func (a *authorizationService) evaluateConstraint(ctx context.Context, params service.CheckRecordAccessParams, permission *resource_model.Permission, now time.Time, userDetails *jwt_model.UserDetails, remainingPropertyCheck *map[string]bool) (bool, []*model.ErrorField) {
	logger := log.WithFields(logging.CtxFields(ctx))

	// check resource permission matches

	if permission.Resource != nil && *permission.Resource != params.Resource.Name {
		// skipping as not related to this resource
		logger.Tracef("Skipping permission as not related to this resource: %v => %v", permission.Resource, params.Resource.Name)
		return false, nil
	}

	if permission.Namespace != nil && *permission.Namespace != params.Resource.Namespace {
		// skipping as not related to this namespace
		logger.Tracef("Skipping permission as not related to this namespace: %v => %v", permission.Namespace, params.Resource.Namespace)
		return false, nil
	}

	if permission.GetRecordIds() != nil {
		var found = false

	mainLoop:
		for _, id := range permission.GetRecordIds() {
			id = a.processValue(id, userDetails)

			if params.Records == nil {
				logger.Tracef("Skipping permission as records not found: %v => %v", permission.RecordIds, id)
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
			logger.Tracef("Skipping permission as record id not matched: %v", permission)
			return false, nil
		}
	}

	if permission.GetOperation() != resource_model.PermissionOperation_FULL && permission.Operation != params.Operation {
		logger.Tracef("Skipping permission as operation not matched: %v", permission)
		return false, nil
	}

	if permission.Before != nil && permission.Before.After(now) {
		logger.Tracef("Skipping permission as before time not matched: %v", permission)
		return false, nil
	}

	if permission.After != nil && permission.Before.After(now) {
		logger.Tracef("Skipping permission as after time not matched: %v", permission)
		return false, nil
	}

	if permission.User != nil && permission.User.Id.String() != userDetails.UserId {
		logger.Tracef("Skipping permission as username not matched: %v", permission)
		return false, nil
	}

	if permission.Property != nil {
		if permission.PropertyMode != nil && *permission.PropertyMode == resource_model.PermissionPropertyMode_PROPERTYMATCHONLY {
			if permission.PropertyValue != nil {
				for _, record := range *params.Records {
					for key := range record.Properties {
						if key == "id" {
							continue
						}
						if key != *permission.Property {
							continue
						}

						(*remainingPropertyCheck)[key] = false

						if permission.PropertyValue != nil {
							var value = *permission.PropertyValue

							value = a.processValue(value, userDetails)

							strActualVal := fmt.Sprintf("%v", record.Properties[*permission.Property].AsInterface())

							if strActualVal != value {
								logger.Tracef("Skipping permission as property value not matched: %v", permission)
								return false, nil
							}
						}
					}
				}
			}
		} else {
			var value = *permission.PropertyValue

			value = a.processValue(value, userDetails)

			if params.Records != nil {
				for _, record := range *params.Records {
					if record.Properties[*permission.Property] == nil {
						logger.Tracef("Skipping permission as property not found: %v", permission)
						return false, nil
					}

					strActualVal := fmt.Sprintf("%v", record.Properties[*permission.Property].AsInterface())

					if strActualVal != value {
						return false, []*model.ErrorField{
							{
								Property: *permission.Property,
								Message:  fmt.Sprintf("Property '%s' is not allowed", *permission.Property),
							},
						}
					}
				}
			}
		}
	}

	return true, nil
}

func (a *authorizationService) processValue(value string, userDetails *jwt_model.UserDetails) string {
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
