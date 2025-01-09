package Validator

import (
	"fmt"
	"strings"
	"testing"
)

// TestIsValidEmail is a test function for IsValidEmail function
func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email         string
		expectedValue bool
		description   string
	}{
		{
			email:         "test@test.com",
			expectedValue: true,
			description:   "Testing with valid email",
		},
		{
			email:         "test",
			expectedValue: false,
			description:   "Testing with invalid email",
		},
		{
			email:         "",
			expectedValue: false,
			description:   "Testing with empty email",
		},
		{
			email:         strings.Repeat("a", 256) + "@test.com",
			expectedValue: false,
			description:   "Testing with email longer than 255 characters",
		},
		{
			email:         strings.Repeat("a", 245) + "@test.com",
			expectedValue: true,
			description:   "Testing with email exactly 255 characters long",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := IsValidEmail(test.email)
			if result != test.expectedValue {
				t.Errorf("Expected %v, but got %v", test.expectedValue, result)
			} else {
				t.Logf("Success: %s", test.description)
			}
		})
	}
}
