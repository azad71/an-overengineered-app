package logger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCallerFuncName(t *testing.T) {
	t.Run("Should return caller function name", func(t *testing.T) {
		src, funcName, err := getCallerFuncSource(1)

		assert.NoError(t, err)
		assert.Equal(t, "TestGetCallerFuncName", funcName)
		assert.Contains(t, src, ":12")
	})

	t.Run("Should return empty string if function depth exceeds", func(t *testing.T) {
		src, funcName, err := getCallerFuncSource(33)

		assert.Error(t, err)
		assert.Equal(t, "", src)
		assert.Equal(t, "", funcName)
	})
}

func TestGetColorForLevel(t *testing.T) {

	tests := []struct {
		level     string
		colorCode string
		colorName string
	}{
		{
			level:     "info",
			colorCode: "\033[0;32m",
			colorName: "green",
		},
		{
			level:     "error",
			colorCode: "\033[1;31m",
			colorName: "red",
		},
		{
			level:     "fatal",
			colorCode: "\033[1;36m",
			colorName: "cyan",
		},
		{
			level:     "panic",
			colorCode: "\033[1;34m",
			colorName: "blue",
		},
		{
			level:     "default",
			colorCode: "\033[0m",
			colorName: "white",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("should return %s for %s level", tt.colorName, tt.level), func(t *testing.T) {
			color := getColorForLevel(tt.level)

			assert.Equal(t, tt.colorCode, color)
		})
	}
}
