package impl

import (
	"context"
	"fmt"
	"github.com/apibrew/apibrew/pkg/errors"
	"github.com/apibrew/apibrew/pkg/helper"
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
	recordInlineSelector *helper.RecordInlineSelector
}

func (a *authorizationService) CheckIsExtensionController(ctx context.Context) errors.ServiceError {
	exp, err := a.CheckRecordAccessWithRecordSelector(ctx, service.CheckRecordAccessParams{
		Resource:  resources.ExtensionResource,
		Operation: resource_model.PermissionOperation_FULL,
	})

	if err != nil {
		return err
	}

	if exp != nil {
		return errors.AccessDeniedError.WithDetails("User must have unconditional access to extension resource")
	}

	return nil
}

func (a *authorizationService) CheckRecordAccess(ctx context.Context, params service.CheckRecordAccessParams) errors.ServiceError {
	_, err := a.CheckRecordAccessWithRecordSelector(ctx, params)

	return err
}

func (a *authorizationService) CheckRecordAccessWithRecordSelector(ctx context.Context, params service.CheckRecordAccessParams) (*resource_model.BooleanExpression, errors.ServiceError) {
	logger := log.WithFields(logging.CtxFields(ctx))

	if util.IsSystemContext(ctx) {
		logger.Trace("System context, skipping authorization check")
		return nil, nil
	}

	if annotations.IsEnabled(params.Resource, annotations.AllowPublicAccess) {
		logger.Trace("Public access is allowed")
		return nil, nil
	}

	if annotations.IsEnabled(params.Resource, a.locatePublicAccessAnnotation(params.Operation)) {
		logger.Trace("Public access is allowed")
		return nil, nil
	}

	userDetails := jwt_model.GetUserDetailsFromContext(ctx)

	if userDetails == nil {
		return nil, errors.AccessDeniedError.WithDetails("Public access is denied to resource: " + params.Resource.Name)
	}

	var permissions []*resource_model.Permission

	permissions = append(permissions, userDetails.Permissions...)

	exp, errorFields := a.evaluateConstraints(ctx, params, permissions, userDetails)

	if len(errorFields) == 0 {
		return exp, nil
	} else {
		return nil, errors.AccessDeniedError.WithDetails("User don't have permission to access this resource").WithErrorFields(errorFields)
	}
}

func (a *authorizationService) locatePublicAccessAnnotation(operation resource_model.PermissionOperation) string {
	switch operation {
	case resource_model.PermissionOperation_READ:
		return annotations.AllowPublicReadAccess
	case resource_model.PermissionOperation_UPDATE:
		return annotations.AllowPublicUpdateAccess
	case resource_model.PermissionOperation_DELETE:
		return annotations.AllowPublicDeleteAccess
	case resource_model.PermissionOperation_CREATE:
		return annotations.AllowPublicCreateAccess
	default:
		return annotations.AllowPublicAccess
	}
}

func (a *authorizationService) evaluateConstraints(ctx context.Context, params service.CheckRecordAccessParams, permissions []*resource_model.Permission, userDetails *jwt_model.UserDetails) (*resource_model.BooleanExpression, []*model.ErrorField) {
	logger := log.WithFields(logging.CtxFields(ctx))

	now := time.Now()
	var errorFields []*model.ErrorField

	var exp *resource_model.BooleanExpression

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
		expLocal, hasAllowFlagLocal := a.evaluateConstraint(ctx, params, permission, now, userDetails)

		logger.Tracef("Constraint evaluation result: %v", hasAllowFlagLocal)

		if hasAllowFlagLocal {
			hasAllowFlag = true
		}

		if expLocal != nil {
			if exp == nil {
				exp = expLocal
			} else {
				exp = &resource_model.BooleanExpression{Or: []resource_model.BooleanExpression{
					*exp, *expLocal,
				}}
			}
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
			Message:  fmt.Sprintf("No permissions matched while accessing resource: %s", params.Resource.Name),
		})
	}

	return exp, errorFields
}

func (a *authorizationService) evaluateConstraint(ctx context.Context, params service.CheckRecordAccessParams, permission *resource_model.Permission, now time.Time, userDetails *jwt_model.UserDetails) (*resource_model.BooleanExpression, bool) {
	logger := log.WithFields(logging.CtxFields(ctx))

	// check resource permission matches

	if permission.Resource != nil && *permission.Resource != params.Resource.Name {
		// skipping as not related to this resource
		logger.Tracef("Skipping permission as not related to this resource: %v => %v", permission.Resource, params.Resource.Name)
		return nil, false
	}

	if permission.Namespace != nil && *permission.Namespace != params.Resource.Namespace {
		// skipping as not related to this namespace
		logger.Tracef("Skipping permission as not related to this namespace: %v => %v", permission.Namespace, params.Resource.Namespace)
		return nil, false
	}

	if permission.GetOperation() != resource_model.PermissionOperation_FULL && permission.Operation != params.Operation {
		logger.Tracef("Skipping permission as operation not matched: %v", permission)
		return nil, false
	}

	if permission.Before != nil && permission.Before.After(now) {
		logger.Tracef("Skipping permission as before time not matched: %v", permission)
		return nil, false
	}

	if permission.After != nil && permission.Before.After(now) {
		logger.Tracef("Skipping permission as after time not matched: %v", permission)
		return nil, false
	}

	if permission.User != nil && permission.User.Id.String() != userDetails.UserId {
		logger.Tracef("Skipping permission as username not matched: %v", permission)
		return nil, false
	}

	if permission.RecordSelector != nil && params.Records != nil {
		var checkedRecords, err = a.recordInlineSelector.SelectRecords(ctx, params.Resource, params.Records, permission.RecordSelector)

		if err != nil {
			logger.Errorf("Error while evaluating record selector: %v", err)
			return nil, false
		}

		if len(checkedRecords) == 0 {
			logger.Tracef("Skipping permission as record selector not matched: %v", permission)
			return nil, false
		}

		*params.Records = checkedRecords
	}

	return permission.RecordSelector, true
}

func NewAuthorizationService() service.AuthorizationService {
	return &authorizationService{
		recordInlineSelector: new(helper.RecordInlineSelector),
	}
}
