package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type testCase struct {
	name     string
	input    int
	expected string
}func TestPyramid(t *testing.T) {

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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var buf bytes.Buffer
			out := pyramid(tc.input)
			fmt.Fprint(&buf, out)

			if buf.String() != tc.expected {
				t.Errorf("Expected output to be %v but got %v", tc.expected, buf.String())
			}
		})
	}
}

