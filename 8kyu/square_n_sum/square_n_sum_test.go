package square_sum

import "testing"

func TestSquareSum(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []int
		expected int
	}{
		{
			desc:     "return 1, [1]",
			input:    []int{1},
			expected: 1,
		},
		{
			desc:     "return 5, [1, 2]",
			input:    []int{1, 2},
			expected: 5,
		},
		{
			desc:     "return 9, [1, 2, 2]",
			input:    []int{1, 2, 2},
			expected: 9,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			actual := SquareSum(tc.input)
			if actual != tc.expected {
				t.Errorf("SquareSum(%v) = %d; expected %d", tc.input, actual, tc.expected)
			}
		})
	}
}
