package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestUtils(t *testing.T) {
	t.Run("IndentJsonString", func(t *testing.T) {
		t.Run("Should return the correct 4-spaces indented JSON string", func(t *testing.T) {
			jsonString := `{"foo":"bar", "baz": 123}`
			indentedJsonString, err := IndentJsonString(jsonString)
			expectedJsonString := "{\n    \"foo\": \"bar\",\n    \"baz\": 123\n}"
			assert.Nil(t, err, fmt.Sprintf("Expected no error, got: '%s'", err))
			assert.Equal(t, expectedJsonString, indentedJsonString)
		})
	})
}
