package errors

import "data-handler/model"

var RecordNotFoundError = newServiceError(model.ErrorCode_RECORD_NOT_FOUND, "record not found")
var ResourceNotFoundError = newServiceError(model.ErrorCode_RESOURCE_NOT_FOUND, "resource not found")
var UnableToLocatePrimaryKey = newServiceError(model.ErrorCode_UNABLE_TO_LOCATE_PRIMARY_KEY, "unable to locate primary key")
var InternalError = newServiceError(model.ErrorCode_INTERNAL_ERROR, "Internal error")
var PropertyNotFoundError = newServiceError(model.ErrorCode_PROPERTY_NOT_FOUND, "Property not found")
var RecordValidationError = newServiceError(model.ErrorCode_RECORD_VALIDATION_ERROR, "Record Validation failed")
var ResourceValidationError = newServiceError(model.ErrorCode_RESOURCE_VALIDATION_ERROR, "Resource Validation failed")
var AuthenticationFailedError = newServiceError(model.ErrorCode_AUTHENTICATION_FAILED, "Authentication failed")
var AccessDeniedError = newServiceError(model.ErrorCode_ACCESS_DENIED, "Access denied")
var BackendConnectionAuthenticationError = newServiceError(model.ErrorCode_BACKEND_ERROR, "Backend error")
var UniqueViolation = newServiceError(model.ErrorCode_UNIQUE_VIOLATION, "Unique violation")
var ForeignKeyViolation = newServiceError(model.ErrorCode_REFERENCE_VIOLATION, "Reference violation")

// RecordValidationError @fixme
var AlreadyExistsError = newServiceError(model.ErrorCode_ALREADY_EXISTS, "Already Exists")
var LogicalError = newServiceError(model.ErrorCode_RECORD_VALIDATION_ERROR, "Logical Error")
