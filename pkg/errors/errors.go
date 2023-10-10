package errors

import (
	"github.com/apibrew/apibrew/pkg/model"
	"google.golang.org/grpc/codes"
)

var RecordNotFoundError = NewServiceError(model.ErrorCode_RECORD_NOT_FOUND, "record not found", codes.NotFound)
var ResourceNotFoundError = NewServiceError(model.ErrorCode_RESOURCE_NOT_FOUND, "resource not found", codes.NotFound)
var UnableToLocatePrimaryKey = NewServiceError(model.ErrorCode_UNABLE_TO_LOCATE_PRIMARY_KEY, "unable to locate primary key", codes.FailedPrecondition)
var InternalError = NewServiceError(model.ErrorCode_INTERNAL_ERROR, "Internal error", codes.Internal)
var LogicalError = NewServiceError(model.ErrorCode_INTERNAL_ERROR, "Logical Error", codes.Internal)
var PropertyNotFoundError = NewServiceError(model.ErrorCode_PROPERTY_NOT_FOUND, "Property not found", codes.FailedPrecondition)
var RecordValidationError = NewServiceError(model.ErrorCode_RECORD_VALIDATION_ERROR, "Record Validation failed", codes.FailedPrecondition)
var ResourceValidationError = NewServiceError(model.ErrorCode_RESOURCE_VALIDATION_ERROR, "resource Validation failed", codes.FailedPrecondition)
var AuthenticationFailedError = NewServiceError(model.ErrorCode_AUTHENTICATION_FAILED, "Authentication failed", codes.Unauthenticated)
var AccessDeniedError = NewServiceError(model.ErrorCode_ACCESS_DENIED, "Access denied", codes.PermissionDenied)
var BackendConnectionAuthenticationError = NewServiceError(model.ErrorCode_BACKEND_ERROR, "Backend error", codes.FailedPrecondition)
var UniqueViolation = NewServiceError(model.ErrorCode_UNIQUE_VIOLATION, "Unique violation", codes.FailedPrecondition)
var ReferenceViolation = NewServiceError(model.ErrorCode_REFERENCE_VIOLATION, "Reference violation", codes.FailedPrecondition)
var UnsupportedOperation = NewServiceError(model.ErrorCode_UNSUPPORTED_OPERATION, "Unsupported Operation", codes.FailedPrecondition)
var ExternalBackendCommunicationError = NewServiceError(model.ErrorCode_EXTERNAL_BACKEND_COMMUNICATION_ERROR, "External Backend communication error", codes.Internal)
var ExternalBackendError = NewServiceError(model.ErrorCode_EXTERNAL_BACKEND_ERROR, "External Backend error", codes.Internal)
var RateLimitError = NewServiceError(model.ErrorCode_RATE_LIMIT_ERROR, "Rate limit exceeded", codes.FailedPrecondition)

// RecordValidationError @fixme
var AlreadyExistsError = NewServiceError(model.ErrorCode_ALREADY_EXISTS, "Already Exists", codes.FailedPrecondition)
