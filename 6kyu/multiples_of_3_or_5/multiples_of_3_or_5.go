package multiples_of_3_or_5

func Multiple3And5(number int) int {
	sumMultiplies := func(m int) int {
		// the better way
		//n := (number - 1) / m
		//return m * n * (n + 1) / 2

		// my first solution
		sum := 0
		for i := (number - 1) / m; i > 0; i-- {
			sum += m * i
		}
		return sum
	}
	return sumMultiplies(3) + sumMultiplies(5) - sumMultiplies(15)
}
