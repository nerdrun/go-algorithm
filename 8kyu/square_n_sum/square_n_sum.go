package square_sum

func SquareSum(numbers []int) int {
	result := 0
	for _, n := range numbers {
		result += n * n
	}
	return result
}
