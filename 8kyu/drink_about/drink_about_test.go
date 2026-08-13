package drink_about

import "testing"

func TestPeopleWithAgeDrink(t *testing.T) {

	tcs := []struct {
		desc     string
		age      int
		expected string
	}{
		{
			desc:     "return drink toddy when 0 age",
			age:      0,
			expected: "drink toddy",
		},
		{
			desc:     "return drink toddy when 13 age",
			age:      13,
			expected: "drink toddy",
		},
		{
			desc:     "return drink toddy when 17 age",
			age:      17,
			expected: "drink coke",
		},
		{
			desc:     "return drink toddy when 18 age",
			age:      18,
			expected: "drink beer",
		},
		{
			desc:     "return drink toddy when 20 age",
			age:      20,
			expected: "drink beer",
		},
		{
			desc:     "return drink toddy when 30 age",
			age:      30,
			expected: "drink whisky",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			result := peopleWithAgeDrink(tc.age)
			if result != tc.expected {
				t.Errorf("peopleWithAgeDrink = %s; want %s", result, tc.expected)
			}
		})
	}
}
