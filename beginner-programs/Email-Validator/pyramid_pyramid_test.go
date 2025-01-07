package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Testpyramid(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Normal Operation Test",
			input:    "3\n",
			expected: "    * \n  * * * \n* * * * * \n",
		},
		{
			name:     "Edge Case Test - Zero Rows",
			input:    "0\n",
			expected: "",
		},
		{
			name:     "Error Handling Test - Negative Number of Rows",
			input:    "-3\n",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			oldStdin := os.Stdin
			oldStdout := os.Stdout
			defer func() {
				os.Stdin = oldStdin
				os.Stdout = oldStdout
			}()

			r, w, _ := os.Pipe()
			os.Stdin = strings.NewReader(tt.input)
			os.Stdout = w

			pyramid()

			w.Close()
			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			if output != tt.expected {
				t.Errorf("got %q, want %q", output, tt.expected)
			}
		})
	}
}
