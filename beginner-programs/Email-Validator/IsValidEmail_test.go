// ********RoostGPT********
/*
Test generated by RoostGPT for test gotest1 using AI Type  and AI Model 

ROOST_METHOD_HASH=IsValidEmail_9ebd397bff
ROOST_METHOD_SIG_HASH=IsValidEmail_d2603fb18f

Scenario 1: Test with a valid email
Details: 
    Description: The test is meant to check whether the function correctly validates a properly formatted email.
Execution: 
    Arrange: Create a string variable with a valid email address.
    Act: Invoke the IsValidEmail function with the valid email.
    Assert: The function should return true.
Validation: 
    The assertion checks whether the return value of the function is true. The expected result is true because the email is valid. This test is important to confirm that the function correctly validates properly formatted emails.

Scenario 2: Test with an invalid email
Details:
    Description: The test is meant to check whether the function correctly invalidates an improperly formatted email.
Execution:
    Arrange: Create a string variable with an invalid email address.
    Act: Invoke the IsValidEmail function with the invalid email.
    Assert: The function should return false.
Validation:
    The assertion checks whether the return value of the function is false. The expected result is false because the email is invalid. This test is important to confirm that the function correctly invalidates improperly formatted emails.

Scenario 3: Test with an empty email
Details: 
    Description: The test is meant to check whether the function correctly invalidates an empty email.
Execution: 
    Arrange: Create a string variable with an empty string.
    Act: Invoke the IsValidEmail function with the empty string.
    Assert: The function should return false.
Validation: 
    The assertion checks whether the return value of the function is false. The expected result is false because an empty string is not a valid email. This test is important to confirm that the function correctly invalidates empty emails.

Scenario 4: Test with a very long email
Details:
    Description: The test is meant to check whether the function correctly invalidates an email that is longer than 255 characters.
Execution:
    Arrange: Create a string variable with an email that is longer than 255 characters.
    Act: Invoke the IsValidEmail function with the long email.
    Assert: The function should return false.
Validation:
    The assertion checks whether the return value of the function is false. The expected result is false because the email is longer than 255 characters. This test is important to confirm that the function correctly invalidates emails that are too long.

Scenario 5: Test with a maximum length valid email
Details:
    Description: The test is meant to check whether the function correctly validates an email that is exactly 255 characters long.
Execution:
    Arrange: Create a string variable with a valid email that is exactly 255 characters long.
    Act: Invoke the IsValidEmail function with the maximum length email.
    Assert: The function should return true.
Validation:
    The assertion checks whether the return value of the function is true. The expected result is true because the email is valid and does not exceed the maximum length. This test is important to confirm that the function correctly validates emails of maximum length.
*/

// ********RoostGPT********
package Validator

import (
    "testing"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

func TestIsValidEmail(t *testing.T) {
    tests := []struct {
        name string
        email string
        want bool
    }{
        {
            name: "Test with a valid email",
            email: "test@example.com",
            want: true,
        },
        {
            name: "Test with an invalid email",
            email: "invalid-email",
            want: false,
        },
        {
            name: "Test with an empty email",
            email: "",
            want: false,
        },
        {
            name: "Test with a very long email",
            email: "thisEmailIsWayTooLongToBeAValidEmailAddressThisEmailIsWayTooLongToBeAValidEmailAddressThisEmailIsWayTooLongToBeAValidEmailAddressThisEmailIsWayTooLongToBeAValidEmailAddressThisEmailIsWayTooLongToBeAValidEmailAddressThisEmailIsWayTooLongToBeAValidEmailAddressThisEmailIsWayTooLongToBeAValidEmailAddressThisEmailIsWayTooLongToBeAValidEmailAddressThisEmailIsWayTooLongToBeAValidEmailAddress@example.com",
            want: false,
        },
        {
            name: "Test with a maximum length valid email",
            email: "thisIsAValidEmailAddressThatIsExactlyTwoHundredAndFiftyFiveCharactersLongThisIsAValidEmailAddressThatIsExactlyTwoHundredAndFiftyFiveCharactersLongThisIsAValidEmailAddressThatIsExactlyTwoHundredAndFiftyFiveCharactersLong@example.com",
            want: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := IsValidEmail(tt.email); got != tt.want {
                t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
            }
        })
    }
}
