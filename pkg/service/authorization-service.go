package service

import (
	"context"
	"github.com/apibrew/apibrew/pkg/abs"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/annotations"
	"github.com/apibrew/apibrew/pkg/service/security"
	"github.com/apibrew/apibrew/pkg/util"
	"time"
)

type authorizationService struct {
}

func (a *authorizationService) CheckRecordAccess(ctx context.Context, params abs.CheckRecordAccessParams) errors.ServiceError {
	if security.IsSystemContext(ctx) {
		return nil
	}

	userDetails := security.GetUserDetailsFromContext(ctx)

	if userDetails == nil && !annotations.IsEnabled(params.Resource, annotations.AllowPublicAccess) {
		return errors.AccessDeniedError.WithDetails("Public access is denied")
	}

	var constraints []*model.SecurityConstraint

	for _, constraint := range params.Resource.SecurityConstraints {
		constraint.Resource = params.Resource.Name
		constraint.Namespace = params.Resource.Namespace
	}

	constraints = append(constraints, params.Resource.SecurityConstraints...)
	constraints = append(constraints, userDetails.SecurityConstraints...)

	result := a.evaluateConstraints(params, constraints, userDetails)

	if result == model.PermitType_PERMIT_TYPE_ALLOW {
		return nil
	} else {
		return errors.AccessDeniedError.WithDetails("User don't have permission to access this resource")
	}
}

func (a *authorizationService) evaluateConstraints(params abs.CheckRecordAccessParams, constraints []*model.SecurityConstraint, userDetails *abs.UserDetails) model.PermitType {
	now := time.Now()

	/*
	  Default policy for checking constraints are like that
	  1. If anyone rejects, then reject
	  2. If none rejects and anyone allows, then allow
	*/

	// Default permit type is disallow
	currentPermitType := model.PermitType_PERMIT_TYPE_REJECT

	for _, constraint := range constraints {
		permit := a.evaluateConstraint(params, constraint, now, userDetails)

		if permit == nil { // constraint does not match our case
			continue
		}

		if *permit == model.PermitType_PERMIT_TYPE_ALLOW {
			currentPermitType = model.PermitType_PERMIT_TYPE_ALLOW
		}

		// if anyone rejects, then reject immediately
		if *permit == model.PermitType_PERMIT_TYPE_REJECT {
			return model.PermitType_PERMIT_TYPE_REJECT
		}
	}

	return currentPermitType
}

func (a *authorizationService) evaluateConstraint(params abs.CheckRecordAccessParams, constraint *model.SecurityConstraint, now time.Time, userDetails *abs.UserDetails) *model.PermitType {
	// check resource constraint matches

	matchPartOkay := a.checkConstraintMatchPart(params, constraint, now)

	if !matchPartOkay {
		return nil
	}

	// we matched our constraint, now check other constraints

	userPartOkay := a.checkConstraintUserPart(constraint, userDetails)
	if !userPartOkay {
		if constraint.RequirePass {
			var permit = model.PermitType_PERMIT_TYPE_REJECT
			return &permit
		} else {
			return nil
		}
	}

	return &constraint.Permit
}

func (a *authorizationService) checkConstraintUserPart(constraint *model.SecurityConstraint, userDetails *abs.UserDetails) bool {
	if constraint.Username != "*" && constraint.Username != "" && constraint.Username != userDetails.Username {
		return false
	}

	if constraint.Role != "*" && constraint.Role != "" && !util.ArrayContains(userDetails.Roles, constraint.Role) {
		return false
	}

	return true
}

func (a *authorizationService) checkConstraintMatchPart(params abs.CheckRecordAccessParams, constraint *model.SecurityConstraint, now time.Time) bool {
	if constraint.Resource != "*" && constraint.Resource != "" && constraint.Resource != params.Resource.Name {
		return false
	}

	if constraint.Namespace != "*" && constraint.Namespace != "" && constraint.Namespace != params.Resource.Namespace {
		return false
	}

	if constraint.RecordIds != nil {
		var found = false

	mainLoop:
		for _, id := range constraint.RecordIds {
			for _, record := range *params.Records {
				if record.Id == id {
					found = true
					break mainLoop
				}
			}
		}

		if !found {
			return false
		}
	}

	if constraint.Operation != model.OperationType_FULL && constraint.Operation != params.Operation {
		return false
	}

	if constraint.Before != nil && constraint.Before.AsTime().After(now) {
		return false
	}

	if constraint.After != nil && constraint.Before.AsTime().After(now) {
		return false
	}

	if constraint.Property != "" && constraint.Property != "*" {
		for _, record := range *params.Records {
			for key := range record.Properties {
				if key == "id" {
					continue
				}
				if key != constraint.Property {
					return false
				}
			}
		}
	}

	return true
}

func NewAuthorizationService() abs.AuthorizationService {
	return &authorizationService{}
}
