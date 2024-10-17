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
*/

// ********RoostGPT********
package Validator

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	// Define test cases
	testCases := []struct {
		email    string
		expected bool
	}{
		{
			email:    "test@example.com",
			expected: true,
		},
		{
			email:    "test.example.com",
			expected: false,
		},
		{
			email:    "test@.com",
			expected: false,
		},
		{
			email:    "test@example..com",
			expected: false,
		},
		{
			email:    "test@",
			expected: false,
		},
		{
			email:    "",
			expected: false,
		},
		{
			email:    "a@b.c",
			expected: true,
		},
		{
			email:    "a@b.co",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {
			// Call the function with the test case input
			output := IsValidEmail(tc.email)

			// Assert that the output is as expected
			if output != tc.expected {
				t.Errorf("IsValidEmail(%v) = %v; want %v", tc.email, output, tc.expected)
			}
		})
	}
}
