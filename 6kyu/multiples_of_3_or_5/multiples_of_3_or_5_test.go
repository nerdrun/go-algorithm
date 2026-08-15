package multiples_of_3_or_5

import "testing"

func TestMultiple3And5(t *testing.T) {
	tcs := []struct {
		desc     string
		input    int
		expected int
	}{
		{
			desc:     "return 0, when input 2",
			input:    2,
			expected: 0,
		},
		{
			desc:     "return 3, when input 3",
			input:    3,
			expected: 0,
		},
		{
			desc:     "return 3, when input 4",
			input:    4,
			expected: 3,
		},
		{
			desc:     "return 8, when input 5",
			input:    5,
			expected: 3,
		},
		{
			desc:     "return 14, when input 6",
			input:    6,
			expected: 8,
		},
		{
			desc:     "return 60, when input 15",
			input:    15,
			expected: 45,
		},
		{
			desc:     "return 225, when input 30",
			input:    30,
			expected: 195,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := Multiple3And5(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %d, got: %d", tc.expected, result)
			}
		})
	}
}
