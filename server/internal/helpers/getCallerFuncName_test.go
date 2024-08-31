package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCallerFuncName(t *testing.T) {
	t.Run("Should return caller function name", func(t *testing.T) {
		callerName := GetCallerFuncName(1)

		assert.Equal(t, "helpers.TestGetCallerFuncName.func1", callerName)
	})

	t.Run("Should return empty string if function depth exceeds", func(t *testing.T) {
		callerName := GetCallerFuncName(33)

		assert.Equal(t, "", callerName)
	})
}
