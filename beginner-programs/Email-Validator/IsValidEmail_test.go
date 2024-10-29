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
    A valid email address should return true when passed to the IsValidEmail function. The importance of this test is to confirm that the function can accurately identify a valid email address.

Scenario 2: Test with an invalid email address

Details:
    Description: The test is meant to check whether the function correctly identifies an invalid email address.
Execution:
    Arrange: Create a string variable that holds an invalid email address.
    Act: Invoke the IsValidEmail function with the invalid email address.
    Assert: Use Go testing facilities to verify that the function returns false.
Validation:
    An invalid email address should return false when passed to the IsValidEmail function. This test is crucial in ensuring that the function can accurately identify invalid email addresses and prevent them from being accepted as valid inputs.

Scenario 3: Test with an email address that exceeds the maximum length

Details:
    Description: The test is meant to check whether the function correctly identifies an email address that exceeds the maximum length.
Execution:
    Arrange: Create a string variable that holds an email address exceeding 255 characters.
    Act: Invoke the IsValidEmail function with the long email address.
    Assert: Use Go testing facilities to verify that the function returns false.
Validation:
    An email address that exceeds the maximum length should return false when passed to the IsValidEmail function. This test is important in ensuring that the function can accurately identify and reject overly long email addresses.

Scenario 4: Test with an empty email address

Details:
    Description: The test is meant to check whether the function correctly identifies an empty email address.
Execution:
    Arrange: Create a string variable that holds an empty string.
    Act: Invoke the IsValidEmail function with the empty string.
    Assert: Use Go testing facilities to verify that the function returns false.
Validation:
    An empty string should return false when passed to the IsValidEmail function as it does not constitute a valid email address. This test is important in ensuring that the function can accurately identify and reject empty email addresses.
*/

// ********RoostGPT********
package Validator

import (
	"testing"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email string
		want  bool
	}{
		{"testEmail@example.com", true},
		{"invalidEmail", false},
		{"overlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddressoverlylongemailaddress@example.com", false},
		{"", false},
	}

	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {
			got := IsValidEmail(tc.email)
			if got != tc.want {
				t.Errorf("IsValidEmail(%q) = %v; want %v", tc.email, got, tc.want)
			}
		})
	}
}
