package Validator

import "testing"

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestIsValidEmail(t *testing.T) {

	testCases := []struct {
		email string
		want  bool
	}{
		{"test@example.com", true},
		{"testexample", false},
		{"test@example.com" + string(make([]byte, 250, 250)), false},
		{"", false},
	}

	for i, tc := range testCases {
		got := IsValidEmail(tc.email)
		if got != tc.want {
			t.Errorf("Test case %d failed: got (%v), want (%v)", i, got, tc.want)
		} else {
			t.Logf("Test case %d success: got (%v)", i, got)
		}
	}
}
