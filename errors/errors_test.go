package errors

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type ParameterContext struct {
	Parameter   string `json:"parameter"`
	Description string `json:"description"`
}

func TestWriteError(t *testing.T) {
	writer := httptest.NewRecorder()

	WriteError(writer, 403, "err_doof", "damsel of oberfürst")

	require.Equal(t, 403, writer.Code)
	require.Equal(t, "application/json", writer.HeaderMap.Get("Content-Type"))
	require.Equal(t, `{"code":"err_doof","description":"damsel of oberfürst"}`+"\n", writer.Body.String())
}

func TestWriteErrorResponse(t *testing.T) {
	writer := httptest.NewRecorder()

	WriteErrorResponse(writer, 403, NewErrorResponse("err_doof", "damsel of oberfürst", nil))

	require.Equal(t, 403, writer.Code)
	require.Equal(t, "application/json", writer.HeaderMap.Get("Content-Type"))
	require.Equal(t, `{"code":"err_doof","description":"damsel of oberfürst"}`+"\n", writer.Body.String())
}

func TestWriteErrorWithContext(t *testing.T) {
	writer := httptest.NewRecorder()

	WriteErrorResponse(writer, 400, NewErrorResponse("err_badparameter", "Some parameters are evil", &ParameterContext{
		Parameter:   "A",
		Description: "A has to be at least 30"}))

	require.Equal(t, 400, writer.Code)
	require.Equal(t, "application/json", writer.HeaderMap.Get("Content-Type"))
	require.Equal(t, `{"code":"err_badparameter","description":"Some parameters are evil","data":{"parameter":"A","description":"A has to be at least 30"}}`+"\n", writer.Body.String())
}

func TestWriteErrorWithArrayContext(t *testing.T) {
	writer := httptest.NewRecorder()

	var parameters [2]*ParameterContext = [...]*ParameterContext{
		&ParameterContext{
			Parameter:   "A",
			Description: "A has to be at least 30"},
		&ParameterContext{
			Parameter:   "B",
			Description: "B has to be at least 10"}}
	WriteErrorResponse(writer, 400, NewErrorResponse("err_badparameter", "Some parameters are evil", parameters))

	require.Equal(t, 400, writer.Code)
	require.Equal(t, "application/json", writer.HeaderMap.Get("Content-Type"))
	require.Equal(t, `{"code":"err_badparameter","description":"Some parameters are evil","data":[{"parameter":"A","description":"A has to be at least 30"},{"parameter":"B","description":"B has to be at least 10"}]}`+"\n", writer.Body.String())
}
