package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// Mock pyramid function to make it testable
func pyramid(rows int) string {
	var k int = 0
	var result string

	for i := 1; i <= rows; i++ {
		k = 0

		for space := 1; space <= rows-i; space++ {
			result += "  "
		}

		for {
			result += "* "
			k++
			if k == 2*i-1 {
				break
			}
		}

		result += "\n"
	}
	return result
}

// Define a test case structure
type testCase struct {
	name     string
	input    int
	expected string
}

func TestPyramid(t *testing.T) {
	// Define test cases
	testCases := []testCase{
		{
			name:     "Normal operation with positive integer input",
			input:    5,
			expected: "        * \n      * * * \n    * * * * * \n  * * * * * * * \n* * * * * * * * * \n",
		},
		{
			name:     "Edge case with input of 0",
			input:    0,
			expected: "",
		},
		{
			name:     "Error handling with negative integer input",
			input:    -5,
			expected: "",
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Capture the output of pyramid
			var buf bytes.Buffer
			out := pyramid(tc.input)
			fmt.Fprint(&buf, out)

			// Compare the output with the expected output
			if buf.String() != tc.expected {
				t.Errorf("Expected output to be %v but got %v", tc.expected, buf.String())
			}
		})
	}
}
