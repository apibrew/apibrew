package service

import (
	"context"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/logging"
	"github.com/apibrew/apibrew/pkg/model"
	"github.com/apibrew/apibrew/pkg/service/security"
	log "github.com/sirupsen/logrus"
	"time"
)

type checkAccessParams struct {
	Resource  *model.Resource
	Records   *[]*model.Record
	Operation model.OperationType
}

func checkAccess(ctx context.Context, params checkAccessParams) errors.ServiceError {
	if security.IsSystemContext(ctx) {
		return nil
	}

	logger := log.WithFields(logging.CtxFields(ctx))

	userDetails := security.GetUserDetailsFromContext(ctx)

	if userDetails == nil {
		return errors.AccessDeniedError.WithDetails("Public access is denied")
	}

	userSecurityContext := userDetails.SecurityContext
	resourceSecurityContext := params.Resource.SecurityContext

	now := time.Now()

	var resourceConstraint *model.SecurityConstraint
	var userConstraint *model.SecurityConstraint

	if resourceSecurityContext != nil {
		for _, constraint := range resourceSecurityContext.Constraints {
			// check resource constraint matches
			matches := true
			matches = matches && (constraint.Principal == "*" || constraint.Principal == "" || constraint.Principal == userDetails.Username)
			matches = matches && constraint.Before.AsTime().Before(now)
			matches = matches && (constraint.After.AsTime().UnixMilli() == 0 || constraint.After.AsTime().After(now))

			if constraint.Operation != model.OperationType_FULL {
				matches = matches && constraint.Operation == params.Operation
			}

			if matches {
				resourceConstraint = constraint
				break
			}
		}
	}

	if userSecurityContext != nil {
		for _, constraint := range userSecurityContext.Constraints {
			// check resource constraint matches
			matches := true
			matches = matches && (constraint.Namespace == "*" || constraint.Namespace == "" || constraint.Namespace == params.Resource.Namespace)
			matches = matches && (constraint.Resource == "*" || constraint.Resource == "" || constraint.Resource == params.Resource.Name)
			matches = matches && constraint.Before.AsTime().Before(now)
			matches = matches && (constraint.After.AsTime().UnixMilli() == 0 || constraint.After.AsTime().After(now))

			if constraint.Operation != model.OperationType_FULL {
				matches = matches && constraint.Operation == params.Operation
			}

			if matches {
				userConstraint = constraint
				break
			}
		}
	}

	logger.Tracef("Security check: params => %v; resourceConstraint => %v; userConstraint => %v", params, resourceConstraint, userConstraint)

	var permit = model.PermitType_PERMIT_TYPE_UNKNOWN

	if resourceConstraint != nil {
		permit = resourceConstraint.Permit
	}

	if permit == model.PermitType_PERMIT_TYPE_ALLOW {
		return nil
	}

	if userConstraint != nil {
		permit = userConstraint.Permit
	}

	if permit == model.PermitType_PERMIT_TYPE_REJECT {
		return errors.AccessDeniedError
	}

	return nil
}
