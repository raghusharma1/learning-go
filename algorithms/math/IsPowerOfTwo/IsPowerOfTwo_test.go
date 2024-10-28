// ********RoostGPT********
/*
Test generated by RoostGPT for test test-golang using AI Type  and AI Model 

ROOST_METHOD_HASH=isPowerOfTwo_fe7a80abf8
ROOST_METHOD_SIG_HASH=isPowerOfTwo_a909b954a6

Scenario 1: Testing with a number that is a power of two

Details:
  Description: This test is meant to check if the function is able to correctly determine if a number is a power of two. The target scenario is when the input number is a power of two (e.g., 2, 4, 8, 16, etc.).

Execution:
  Arrange: No arrangement is necessary for this test.
  Act: Invoke the isPowerOfTwo function with a number that is a power of two.
  Assert: Use Go's testing facilities to verify that the function returns true.

Validation:
  The assertion for this test is simple, as we expect the function to return true when a power of two is provided. This test is important as it validates the function's ability to correctly identify powers of two.

Scenario 2: Testing with a number that is not a power of two

Details:
  Description: This test is meant to check if the function can correctly identify numbers that are not powers of two. The target scenario is when the input number is not a power of two (e.g., 3, 5, 6, 7, etc.).

Execution:
  Arrange: No arrangement is necessary for this test.
  Act: Invoke the isPowerOfTwo function with a number that is not a power of two.
  Assert: Use Go's testing facilities to verify that the function returns false.

Validation:
  The assertion for this test is straightforward, as we expect the function to return false when the input number is not a power of two. This test is important as it validates the function's ability to correctly identify non-powers of two.

Scenario 3: Testing with a number less than two

Details:
  Description: This test is meant to check if the function correctly handles numbers less than two. The target scenario is when the input number is less than two (e.g., 0, 1, -1, -2, etc.).

Execution:
  Arrange: No arrangement is necessary for this test.
  Act: Invoke the isPowerOfTwo function with a number that is less than two.
  Assert: Use Go's testing facilities to verify that the function returns false.

Validation:
  The assertion for this test is straightforward, as we expect the function to return false when the input number is less than two. This test is important as it validates the function's ability to correctly handle edge cases.
  
Scenario 4: Testing with a large power of two

Details:
  Description: This test is meant to check if the function can correctly handle large numbers that are powers of two. The target scenario is when the input number is a large power of two (e.g., 2^30).

Execution:
  Arrange: No arrangement is necessary for this test.
  Act: Invoke the isPowerOfTwo function with a large number that is a power of two.
  Assert: Use Go's testing facilities to verify that the function returns true.

Validation:
  The assertion for this test is straightforward, as we expect the function to return true when the input number is a large power of two. This test is important as it validates the function's ability to handle large numbers correctly.
*/

// ********RoostGPT********
package IsPowerOfTwo

import (
    "testing"
)

// TestIsPowerOfTwo is a table driven test for checking if given number is power of two
func TestIsPowerOfTwo(t *testing.T) {
    tests := []struct {
        name string
        num  int
        want bool
    }{
        {
            name: "Testing with a number that is a power of two",
            num:  16,
            want: true,
        },
        {
            name: "Testing with a number that is not a power of two",
            num:  6,
            want: false,
        },
        {
            name: "Testing with a number less than two",
            num:  1,
            want: false,
        },
        {
            name: "Testing with a large power of two",
            num:  1 << 30,
            want: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := isPowerOfTwo(tt.num); got != tt.want {
                t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
            }
        })
    }
}
