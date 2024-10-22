// ********RoostGPT********
/*
Test generated by RoostGPT for test gotest1 using AI Type  and AI Model 

ROOST_METHOD_HASH=IsValidEmail_9ebd397bff
ROOST_METHOD_SIG_HASH=IsValidEmail_d2603fb18f

Scenario 1: Test with a valid email address

Details:
    Description: The test is meant to check whether the function correctly validates a properly formatted email address.
Execution:
    Arrange: Create a string variable that holds a valid email address.
    Act: Invoke the IsValidEmail function with the valid email address.
    Assert: Use Go testing facilities to verify that the function returns true.
Validation:
    A valid email address should return true when passed to the IsValidEmail function. The importance of this test is to confirm that the function correctly validates properly formatted email addresses.

Scenario 2: Test with an invalid email address

Details:
    Description: The test is meant to check whether the function correctly invalidates an improperly formatted email address.
Execution:
    Arrange: Create a string variable that holds an invalid email address.
    Act: Invoke the IsValidEmail function with the invalid email address.
    Assert: Use Go testing facilities to verify that the function returns false.
Validation:
    An invalid email address should return false when passed to the IsValidEmail function. This test is important to confirm that the function correctly invalidates improperly formatted email addresses.

Scenario 3: Test with a string that exceeds 255 characters

Details:
    Description: The test is meant to check whether the function correctly invalidates an email address that exceeds 255 characters.
Execution:
    Arrange: Create a string variable that holds an email address with more than 255 characters.
    Act: Invoke the IsValidEmail function with the email address.
    Assert: Use Go testing facilities to verify that the function returns false.
Validation:
    An email address that exceeds 255 characters should return false when passed to the IsValidEmail function. This test is important to ensure that the function correctly invalidates email addresses that are too long.

Scenario 4: Test with an empty string

Details:
    Description: The test is meant to check whether the function correctly invalidates an empty string.
Execution:
    Arrange: Create a string variable that holds an empty string.
    Act: Invoke the IsValidEmail function with the empty string.
    Assert: Use Go testing facilities to verify that the function returns false.
Validation:
    An empty string should return false when passed to the IsValidEmail function. This test is important to ensure that the function correctly invalidates empty strings.

Scenario 5: Test with a string that is exactly 255 characters long

Details:
    Description: The test is meant to check whether the function correctly validates an email address that is exactly 255 characters long.
Execution:
    Arrange: Create a string variable that holds an email address with exactly 255 characters.
    Act: Invoke the IsValidEmail function with the email address.
    Assert: Use Go testing facilities to verify that the function returns true.
Validation:
    An email address that is exactly 255 characters long should return true when passed to the IsValidEmail function. This test is important to ensure that the function correctly validates email addresses that are at the maximum allowed length.
*/

// ********RoostGPT********
package Validator

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"test@example.com", true},
		{"test.example.com", false},
		{createString(256), false},
		{"", false},
		{createString(255), true},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := IsValidEmail(tt.input)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

// createString is a helper function to create a string of n length
func createString(n int) string {
	var str string
	for i := 0; i < n; i++ {
		str += "a"
	}
	return str + "@example.com"
}
