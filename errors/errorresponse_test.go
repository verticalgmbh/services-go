package errors

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCode(t *testing.T) {
	err := NewErrorResponse("err_4464", "", nil)

	require.Equal(t, "err_4464", err.Code)
}

func TestDescription(t *testing.T) {
	err := NewErrorResponse("", "Lulu lala land", nil)

	require.Equal(t, "Lulu lala land", err.Description)
}

func TestError(t *testing.T) {
	err := NewErrorResponse("", "Lulu lala land", nil)

	require.Equal(t, "Lulu lala land", err.Error())
}

func TestContext(t *testing.T) {
	err := NewErrorResponse("", "", "blubb")

	require.Equal(t, "blubb", err.Data)
}
