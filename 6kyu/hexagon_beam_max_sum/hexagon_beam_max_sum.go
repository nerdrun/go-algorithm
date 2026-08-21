package hexagon_beam_max_sum

func MaxHexagonBeam(n int, seq []int) int {
	total := 0
	j := 0

	longestLength := n + (n - 1)
	for i := 0; i < longestLength; i++ {
		total += seq[j]
		j++
		if j == len(seq) {
			j = 0
		}
	}
	return total
	// 2 + 3 + 4 + 2
	//   2 3 4
	//  2 3 4 2
	// 3 4 2 3 4
	//  2 3 4 2
	//   2 3 4
}
