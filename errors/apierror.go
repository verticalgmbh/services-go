package errors

// APIError error returned by an api service
type APIError struct {
	StatusCode int            // http status code
	Response   *ErrorResponse // error response to return as body
}

// Error error func to satisfy error interface
//       returns description of wrapped error
func (err *APIError) Error() string {
	return err.Response.Description
}

// NewAPIError creates a api error
//
// **Parameters**
//   - statuscode : http status code
//   - code       : error code uniquely identifying error type
//   - description: error description
//   - data       : error context data
//
// **Returns**
//   *APIError: created api error
func NewAPIError(statuscode int, code string, description string, data interface{}) *APIError {
	return &APIError{
		StatusCode: statuscode,
		Response: &ErrorResponse{
			Code:        code,
			Description: description,
			Data:        data}}
}
