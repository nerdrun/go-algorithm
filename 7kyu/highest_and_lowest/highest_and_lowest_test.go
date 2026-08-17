package highest_and_lowest

import "testing"

func TestHighAndLow(t *testing.T) {
	tcs := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "return 1 for 1",
			input:    "1",
			expected: "1 1",
		},
		{
			desc:     "return 2 1 for 1 2",
			input:    "1 2",
			expected: "2 1",
		},
		{
			desc:     "return 3 1 for 1 2 3",
			input:    "1 2 3",
			expected: "3 1",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := HighAndLow(tc.input)
			if result != tc.expected {
				t.Errorf("HighAndLow(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}
