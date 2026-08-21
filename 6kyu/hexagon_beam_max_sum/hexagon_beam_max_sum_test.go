package hexagon_beam_max_sum

import "testing"

func TestMaxHexagonBeam(t *testing.T) {
	tcs := []struct {
		desc     string
		n        int
		seq      []int
		expected int
	}{
		{
			desc:     "return 1 when n = 1, seq = [1]",
			n:        1,
			seq:      []int{1},
			expected: 1,
		},
		{
			desc:     "return 2 when n = 1, seq = [2]",
			n:        1,
			seq:      []int{2},
			expected: 2,
		},
		{
			desc:     "return 4 when n = 2, seq = [1, 2]",
			n:        2,
			seq:      []int{1, 2},
			expected: 4,
		},
		{
			desc:     "return 16 when n = 3, seq = [2, 3, 4]",
			n:        3,
			seq:      []int{2, 3, 4},
			expected: 16,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := MaxHexagonBeam(tc.n, tc.seq)
			if result != tc.expected {
				t.Errorf("got %d, want %d", result, tc.expected)
			}
		})
	}
}
