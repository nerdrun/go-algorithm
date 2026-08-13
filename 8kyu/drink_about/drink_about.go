package drink_about

func peopleWithAgeDrink(age int) string {
	result := "whisky"
	if age < 14 {
		result = "toddy"
	} else if age < 18 {
		result = "coke"
	} else if age < 21 {
		result = "beer"
	}
	return "drink " + result
}
