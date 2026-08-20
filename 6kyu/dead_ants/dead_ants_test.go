package dead_ants

import "testing"

func TestDeadAntCount(t *testing.T) {
	tcs := []struct {
		desc     string
		input    string
		expected int
	}{
		{
			desc:     "0 death",
			input:    "ant",
			expected: 0,
		},
		{
			desc:     "0 death",
			input:    "...ant",
			expected: 0,
		},
		{
			desc:     "1 death",
			input:    "...aant",
			expected: 1,
		},
		{
			desc:     "2 deaths",
			input:    "...aanta",
			expected: 2,
		},
		{
			desc:     "1 deaths",
			input:    "...aantn",
			expected: 1,
		},
		{
			desc:     "2 deaths",
			input:    "...aantna",
			expected: 2,
		},
		{
			desc:     "3 deaths",
			input:    "...ant...ant..nat.ant.t..ant...ant..ant..ant.anant..t",
			expected: 3,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := DeadAntCount(tc.input)
			if result != tc.expected {
				t.Errorf("result: %d, expected: %d", result, tc.expected)
			}
		})
	}
}
