// ********RoostGPT********
/*
Test generated by RoostGPT for test test-golang using AI Type  and AI Model 

ROOST_METHOD_HASH=fibonacci_5d6017f964
ROOST_METHOD_SIG_HASH=fibonacci_3c2494e9fa

Scenario 1: Test with a positive integer

Details:
  Description: This test is designed to verify that the fibonacci function returns the correct Fibonacci series value for a given positive integer. The Fibonacci series is a sequence of numbers in which each number is the sum of the two preceding ones, usually starting with 0 and 1. 

Execution:
  Arrange: No arrangement is required as we are directly feeding the function with raw integers.
  Act: Invoke the fibonacci function with a positive integer, for instance, 5.
  Assert: Verify that the returned value is the correct Fibonacci number at the 5th position in the series.

Validation:
  The assertion checks that the function correctly calculates the Fibonacci series value for a given positive integer. It's important to validate this functionality as it's the core functionality of the function.

Scenario 2: Test with zero

Details:
  Description: This test checks if the function correctly handles the edge case where the input is zero. According to the Fibonacci series, the first number (at position zero) should be 0.

Execution:
  Arrange: No arrangement is required as we are directly feeding the function with raw integers.
  Act: Invoke the fibonacci function with zero as an argument.
  Assert: Verify that the returned value is 0.

Validation:
  The assertion checks that the function correctly handles the edge case of zero as input. This is important to ensure that the function doesn't break or produce incorrect results when encountering edge cases.

Scenario 3: Test with negative integer

Details:
  Description: This test is designed to check how the function handles negative integers. The Fibonacci series is not defined for negative integers, so the function should ideally return an error or a predefined value.

Execution:
  Arrange: No arrangement is required as we are directly feeding the function with raw integers.
  Act: Invoke the fibonacci function with a negative integer, for instance, -5.
  Assert: Verify that the returned value is either an error or a predefined value (for instance, -1).

Validation:
  The assertion checks that the function correctly handles negative integers. This is important to ensure that the function behaves predictably even when given invalid input, and doesn't break or produce incorrect results.
*/

// ********RoostGPT********
package Fibonacci

import (
	"testing"
)

func TestFibonacci399(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{5, 5}, // Scenario 1: Test with a positive integer
		{0, 0}, // Scenario 2: Test with zero
		{-5, -1}, // Scenario 3: Test with negative integer
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			ans := fibonacci(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
