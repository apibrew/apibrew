package errors

import "data-handler/stub/model"

var NotFoundError ServiceError = newServiceError(model.ErrorCode_RECORD_NOT_FOUND, "record not found")
var UnableToLocatePrimaryKey ServiceError = newServiceError(model.ErrorCode_UNABLE_TO_LOCATE_PRIMARY_KEY, "unable to locate primary key")
var InternalError ServiceError = newServiceError(model.ErrorCode_INTERNAL_ERROR, "Internal error")
var PropertyNotFoundError ServiceError = newServiceError(model.ErrorCode_PROPERTY_NOT_FOUND, "Property not found")
var RecordValidationError ServiceError = newServiceError(model.ErrorCode_RECORD_VALIDATION_ERROR, "Validation failed")
var AuthenticationFailedError ServiceError = newServiceError(model.ErrorCode_AUTHENTICATION_FAILED, "Authentication failed")
var AccessDeniedError ServiceError = newServiceError(model.ErrorCode_AUTHENTICATION_FAILED, "Authentication failed")
var BackendConnectionAuthenticationError ServiceError = newServiceError(model.ErrorCode_AUTHENTICATION_FAILED, "Authentication failed")
var UniqueViolation ServiceError = newServiceError(model.ErrorCode_AUTHENTICATION_FAILED, "Authentication failed")

// RecordValidationError @fixme
var AlreadyExistsError ServiceError = newServiceError(model.ErrorCode_RECORD_VALIDATION_ERROR, "Validation failed")
var LogicalError ServiceError = newServiceError(model.ErrorCode_RECORD_VALIDATION_ERROR, "Validation failed")
