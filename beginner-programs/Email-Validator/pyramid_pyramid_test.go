package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// Refactor the pyramid function to make it testable
// It now takes the number of rows as a parameter and returns the pyramid as a string
func pyramid(rows int) string {
	var pyramidBuilder strings.Builder
	var k int = 0

	for i := 1; i <= rows; i++ {
		k = 0

		for space := 1; space <= rows-i; space++ {
			pyramidBuilder.WriteString("  ")
		}

		for {
			pyramidBuilder.WriteString("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}

		pyramidBuilder.WriteString("\n")
	}

	return pyramidBuilder.String()
}

func TestPyramid(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name   string
		rows   int
		expect string
	}{
		{
			name: "Valid Pyramid Generation",
			rows: 3,
			expect: "    * \n  * * * \n* * * * * \n",
		},
		{
			name:   "Pyramid with Zero Rows",
			rows:   0,
			expect: "",
		},
		{
			name:   "Pyramid with Negative Rows",
			rows:   -1,
			expect: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act: Invoke the pyramid function with the number of rows from the test case
			result := pyramid(tc.rows)

			// Assert: Check that the result matches the expected output
			if result != tc.expect {
				t.Errorf("Expected pyramid:\n%s\nbut got:\n%s\n", tc.expect, result)
			} else {
				t.Logf("Success: %s\n", tc.name)
			}
		})
	}
}
