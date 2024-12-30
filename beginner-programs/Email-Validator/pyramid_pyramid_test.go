package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func Testpyramid(t *testing.T) {

	testCases := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Normal Operation Test",
			input:       "3\n",
			expected:    "    * \n  * * * \n* * * * * \n",
			expectError: false,
		},
		{
			name:        "Edge Case Test - Zero Rows",
			input:       "0\n",
			expected:    "",
			expectError: false,
		},
		{
			name:        "Error Handling Test - Negative Number of Rows",
			input:       "-1\n",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			os.Stdin = strings.NewReader(tc.input)

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			pyramid()

			w.Close()
			os.Stdout = oldStdout

			out, _ := io.ReadAll(r)

			if string(out) != tc.expected {
				t.Errorf("unexpected output, got: %v, want: %v", string(out), tc.expected)
			}

			if tc.expectError {
				if string(out) != "" {
					t.Errorf("expected an error but got output: %v", string(out))
				}
			}
		})
	}
}
