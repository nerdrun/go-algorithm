package keep_hydrated

import "testing"

func TestLitres(t *testing.T) {
	tcs := []struct {
		desc     string
		time     float64
		expected int
	}{
		{
			desc:     "return litres 1 for time 3",
			time:     3,
			expected: 1,
		},
		{
			desc:     "return litres 2 for time 4",
			time:     4,
			expected: 2,
		},
		{
			desc:     "return litres 3 for time 6.7",
			time:     6.7,
			expected: 3,
		},
		{
			desc:     "return litres 5 for time 11.8",
			time:     11.8,
			expected: 5,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := Litres(tc.time)
			if result != tc.expected {
				t.Errorf("Litres(%f) => %d, want %d", tc.time, result, tc.expected)
			}
		})
	}
}
