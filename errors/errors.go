package errors

import (
	"encoding/json"
	"net/http"
)

// WriteAPIError writes an api error type
//
// **Parameters**
//   - writer     : writer used to write error to
//   - error      : error containing error data
func WriteAPIError(writer http.ResponseWriter, error *APIError) {
	WriteErrorResponse(writer, error.StatusCode, error.Response)
}

// WriteErrorResponse writes an error response type
//
// **Parameters**
//   - writer     : writer used to write error to
//   - httpcode   : httpcode to write to header
//   - error      : error containing error data
func WriteErrorResponse(writer http.ResponseWriter, httpcode int, error *ErrorResponse) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpcode)

	encoder := json.NewEncoder(writer)
	encoder.Encode(error)
}

// WriteError writes an error as an http response
//
// **Parameters**
//   - writer     : writer used to write error to
//   - httpcode   : httpcode to write to header
//   - errorcode  : errorcode to use
//   - description: description for error
func WriteError(writer http.ResponseWriter, httpcode int, errorcode string, description string) {
	WriteErrorResponse(writer, httpcode, &ErrorResponse{
		Code:        errorcode,
		Description: description})
}

// WriteErrorWithContext writes an error as an http response adding context information
//
// **Parameters**
//   - writer     : writer used to write error to
//   - httpcode   : httpcode to write to header
//   - errorcode  : errorcode to use
//   - description: description for error
//   - data       : context data
func WriteErrorWithContext(writer http.ResponseWriter, httpcode int, errorcode string, description string, data interface{}) {
	WriteErrorResponse(writer, httpcode, &ErrorResponse{
		Code:        errorcode,
		Description: description,
		Data:        data})
}
