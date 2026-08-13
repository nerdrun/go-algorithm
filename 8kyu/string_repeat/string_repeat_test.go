package string_repeat

import "testing"

func TestRepeatStr(t *testing.T) {
	tcs := []struct {
		desc        string
		repetitions int
		value       string
		expected    string
	}{
		{
			desc:        "1 repetition",
			repetitions: 1,
			value:       "abc",
			expected:    "abc",
		},
		{
			desc:        "2 repetition",
			repetitions: 2,
			value:       "Gianni",
			expected:    "GianniGianni",
		},
		{
			desc:        "7 repetition",
			repetitions: 7,
			value:       "I",
			expected:    "IIIIIII",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := RepeatStr(tc.repetitions, tc.value)
			if result != tc.expected {
				t.Errorf("RepeatStr = %s; want %s", result, tc.expected)
			}
		})
	}
}
