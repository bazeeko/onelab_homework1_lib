package onelab_homework1_lib

import (
	"testing"
)

func TestChangeLetterCase(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"alO", "ALo"},
		{"laLAlaA", "LAlaLAa"},
		{"aAaaaAAaaA", "AaAAAaaAAa"},
	}
	// testCase := []string{"alO", "laLAlaA", "aAaaaAAaaA"}
	// testResult := []string{"ALo", "LAlaLAa", "AaAAAaaAAa"}

	for _, tc := range testCases {
		if str := ChangeLetterCase(tc.input); str != tc.output {
			t.Errorf("ChangeLetterCase(%s) => got: %s, want: %s", tc.input, str, tc.output)
		}
	}
}
