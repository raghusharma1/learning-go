package Validator

import (
	"testing"
)

// TestIsValidEmail is a test function for the IsValidEmail function
func TestIsValidEmail(t *testing.T) {
	// Define test cases
	testCases := []struct {
		email string
		want  bool
	}{
		{"test@example.com", true}, // Scenario 1: Testing with a valid email address
		{"testexample", false},     // Scenario 2: Testing with an invalid email address
		{"test@example.com" + string(make([]byte, 250, 250)), false}, // Scenario 3: Testing with an email address that exceeds 255 characters
		{"", false}, // Scenario 4: Testing with an empty email address
	}

	// Execute test cases
	for i, tc := range testCases {
		got := IsValidEmail(tc.email)
		if got != tc.want {
			t.Errorf("Test case %d failed: got (%v), want (%v)", i, got, tc.want)
		} else {
			t.Logf("Test case %d success: got (%v)", i, got)
		}
	}
}
