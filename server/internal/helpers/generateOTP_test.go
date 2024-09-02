package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOTP(t *testing.T) {

	t.Run("GenerateOTP should return 6 digit otp", func(t *testing.T) {
		maxDigits := int8(6)
		otp, err := GenerateOTP(maxDigits)

		assert.NoError(t, err)

		assert.Len(t, otp, int(maxDigits))
	})

	t.Run("GenerateOTP should return empty string and error if maxDigit exceeds", func(t *testing.T) {

		maxDigits := int8(11)

		otp, err := GenerateOTP(maxDigits)

		assert.Error(t, err)
		assert.Equal(t, "", otp)
	})
}
