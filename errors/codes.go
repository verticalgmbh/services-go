package errors

// this package provides default error codes for services

// ErrorCodeRequestMissingParameter the request is missing some parameter
const ErrorCodeRequestMissingParameter = "request_missing_parameter"

// ErrorCodeRequestBadFormat some request is not in the expected format
const ErrorCodeRequestBadFormat = "request_bad_format"

// ErrorCodeDataNotFound a requested entity or object was not found
const ErrorCodeDataNotFound = "data_not_found"

// ErrorCodeInternal an internal server error occured
const ErrorCodeInternal = "internal_server_error"
