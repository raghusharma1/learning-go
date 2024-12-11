package Validator

import (
	"testing"
)

// TestIsValidEmail is a test function for IsValidEmail function
func TestIsValidEmail(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "Valid Email Test",
			email:    "test@example.com",
			expected: true,
		},
		{
			name:     "Invalid Email Test",
			email:    "test.example.com",
			expected: false,
		},
		{
			name:     "Email Length Exceeds Maximum Limit Test",
			email:    "a@b.c" + string(make([]byte, 255)),
			expected: false,
		},
		{
			name:     "Empty Email Test",
			email:    "",
			expected: false,
		},
		{
			name:     "Email with Valid Length but Invalid Format Test",
			email:    "test@.com",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := IsValidEmail(tc.email); result != tc.expected {
				t.Fatalf("Failed %s: Expected %v but got %v", tc.name, tc.expected, result)
			} else {
				t.Logf("Success %s", tc.name)
			}
		})
	}
}
