package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello json"
	result := jsonError(msg)
	require.Equal(t, string([]byte(`{"message":"Hello json"}`)), string(result))
}
