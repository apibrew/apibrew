package errors

import (
	"github.com/tislib/data-handler/pkg/model"
	"google.golang.org/grpc/codes"
)

var RecordNotFoundError = newServiceError(model.ErrorCode_RECORD_NOT_FOUND, "record not found", codes.NotFound)
var ResourceNotFoundError = newServiceError(model.ErrorCode_RESOURCE_NOT_FOUND, "resource not found", codes.NotFound)
var UnableToLocatePrimaryKey = newServiceError(model.ErrorCode_UNABLE_TO_LOCATE_PRIMARY_KEY, "unable to locate primary key", codes.FailedPrecondition)
var InternalError = newServiceError(model.ErrorCode_INTERNAL_ERROR, "Internal error", codes.Internal)
var PropertyNotFoundError = newServiceError(model.ErrorCode_PROPERTY_NOT_FOUND, "Property not found", codes.FailedPrecondition)
var RecordValidationError = newServiceError(model.ErrorCode_RECORD_VALIDATION_ERROR, "Record Validation failed", codes.FailedPrecondition)
var ResourceValidationError = newServiceError(model.ErrorCode_RESOURCE_VALIDATION_ERROR, "resource Validation failed", codes.FailedPrecondition)
var AuthenticationFailedError = newServiceError(model.ErrorCode_AUTHENTICATION_FAILED, "Authentication failed", codes.Unauthenticated)
var AccessDeniedError = newServiceError(model.ErrorCode_ACCESS_DENIED, "Access denied", codes.PermissionDenied)
var BackendConnectionAuthenticationError = newServiceError(model.ErrorCode_BACKEND_ERROR, "Backend error", codes.FailedPrecondition)
var UniqueViolation = newServiceError(model.ErrorCode_UNIQUE_VIOLATION, "Unique violation", codes.FailedPrecondition)
var ReferenceViolation = newServiceError(model.ErrorCode_REFERENCE_VIOLATION, "Reference violation", codes.FailedPrecondition)

// RecordValidationError @fixme
var AlreadyExistsError = newServiceError(model.ErrorCode_ALREADY_EXISTS, "Already Exists", codes.FailedPrecondition)
var LogicalError = newServiceError(model.ErrorCode_RECORD_VALIDATION_ERROR, "Logical Error", codes.FailedPrecondition)
