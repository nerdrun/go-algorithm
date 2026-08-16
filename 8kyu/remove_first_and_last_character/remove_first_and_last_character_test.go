package remove_first_and_last_character

import "testing"

func TestRemoveChar(t *testing.T) {
	tcs := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "return empty string for two characters",
			input:    "ab",
			expected: "",
		},
		{
			desc:     "return b for abc",
			input:    "abc",
			expected: "b",
		},
		{
			desc:     "return bc for abcd",
			input:    "abcd",
			expected: "bc",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := RemoveChar(tc.input)
			if result != tc.expected {
				t.Errorf("RemoveChar(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}
