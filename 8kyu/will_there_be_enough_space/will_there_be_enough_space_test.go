package will_there_be_enough_space

import "testing"

func TestEnough(t *testing.T) {
	tcs := []struct {
		desc     string
		cap      int
		on       int
		wait     int
		expected int
	}{
		{
			desc:     "expected 0 cap 1, on 0, wait 1",
			cap:      1,
			on:       0,
			wait:     1,
			expected: 0,
		},
		{
			desc:     "expected 1 cap 1, on 1, wait 1",
			cap:      1,
			on:       1,
			wait:     1,
			expected: 1,
		},
		{
			desc:     "expected 1 cap 10, on 6, wait 5",
			cap:      10,
			on:       6,
			wait:     5,
			expected: 1,
		},
		{
			desc:     "expected 1 cap 10, on 5, wait 6",
			cap:      10,
			on:       5,
			wait:     6,
			expected: 1,
		},
		{
			desc:     "expected 1 cap 100, on 50, wait 80",
			cap:      100,
			on:       50,
			wait:     80,
			expected: 30,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := Enough(tc.cap, tc.on, tc.wait)
			if result != tc.expected {
				t.Errorf("got %d, want %d", result, tc.expected)
			}
		})
	}
}
