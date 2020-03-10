package errors

// ErrorResponse error in api context
type ErrorResponse struct {
	Code        string      `json:"code"`           // Code identifying type of error
	Description string      `json:"description"`    // detailed error message
	Data        interface{} `json:"data,omitempty"` // optional additional context data for error
}

// NewErrorResponse creates a new error response
//
// **Parameters**
//   - code       : error code uniquely identifying error type
//   - description: error description
//   - data       : error context data
//
// **Returns**
//   *ErrorResponse: created error response
func NewErrorResponse(code string, description string, data interface{}) *ErrorResponse {
	return &ErrorResponse{
		Code:        code,
		Description: description,
		Data:        data}
}

// WrapAPIError wraps an existing error to a error response
//              this automatically specified an internal server error as code
//
// **Parameters**
//   - err: error to wrap
//
// **Returns**
//   *ErrorResponse: created error response
func WrapAPIError(statuscode int, err error) *APIError {
	return &APIError{
		StatusCode: statuscode,
		Response: &ErrorResponse{
			Code:        ErrorCodeInternal,
			Description: err.Error()}}
}

// Error implementation of error interface
//
// **Returns**
//   - string: error description
func (err *ErrorResponse) Error() string {
	return err.Description
}
