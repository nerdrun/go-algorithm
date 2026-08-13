package is_he_gonna_survive

import "testing"

func TestHero(t *testing.T) {
	tcs := []struct {
		desc     string
		bullets  int
		dragons  int
		expected bool
	}{
		{
			desc:     "success when bullets 2, dragon 1",
			bullets:  2,
			dragons:  1,
			expected: true,
		},
		{
			desc:     "success when bullets 3, dragon 1",
			bullets:  3,
			dragons:  1,
			expected: true,
		},
		{
			desc:     "success when bullets 3, dragon 2",
			bullets:  3,
			dragons:  2,
			expected: false,
		},
		{
			desc:     "success when bullets 42, dragon 22",
			bullets:  42,
			dragons:  22,
			expected: false,
		},
		{
			desc:     "success when bullets 42, dragon 21",
			bullets:  42,
			dragons:  21,
			expected: true,
		},
		{
			desc:     "success when bullets 42, dragon 20",
			bullets:  42,
			dragons:  20,
			expected: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := Hero(tc.bullets, tc.dragons)
			if result != tc.expected {
				t.Errorf("expected: %t, got: %t", tc.expected, result)
			}
		})
	}
}
