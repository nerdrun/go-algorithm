package disemvowel_trolls

import "testing"

func TestDisemvowel(t *testing.T) {
	tcs := []struct {
		desc     string
		comment  string
		expected string
	}{
		{
			desc:     "return empty for a",
			comment:  "a",
			expected: "",
		},
		{
			desc:     "return empty for ae",
			comment:  "ae",
			expected: "",
		},
		{
			desc:     "return empty for aei",
			comment:  "aei",
			expected: "",
		},
		{
			desc:     "return empty for aeiou",
			comment:  "aeiou",
			expected: "",
		},
		{
			desc:     "return Ths wbst s fr lsrs LL! for This website is for losers LOL!",
			comment:  "This website is for losers LOL!",
			expected: "Ths wbst s fr lsrs LL!",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := Disemvowel(tc.comment)
			if got != tc.expected {
				t.Errorf("Disemvowel(%q) got %q, want %q", tc.comment, got, tc.expected)
			}
		})
	}
}
