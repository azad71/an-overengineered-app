package helpers

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestIsDateBefore(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("isdatebefore", IsDateBefore)

	dateItems := []struct {
		have string
		want bool
	}{
		{"2010-01-12", true},  // Valid date more than 12 years ago
		{"2013-01-12", false}, // Invalid date within the last 12 years
		{time.Now().AddDate(-12, 0, 0).Format("2006-01-02"), false}, // Invalid date exactly 12 years ago
		{"13-01-2000", false}, // Invalid date format
	}

	for _, item := range dateItems {
		err := validate.Var(item.have, "isdatebefore")
		valid := err == nil
		assert.Equal(t, item.want, valid, "Test case failed for input: %v", item.have)
	}
}
