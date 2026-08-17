package highest_and_lowest

import (
	"fmt"
	"math"
	"strconv"
)

func HighAndLow(in string) string {
	highest := math.MinInt
	lowest := math.MaxInt
	start := 0
	n := len(in)

	for i := 0; i <= n; i++ {
		if i == n || in[i] == ' ' {
			if start < i {
				val, _ := strconv.Atoi(in[start:i])
				if val > highest {
					highest = val
				}
				if val < lowest {
					lowest = val
				}
			}
			start = i + 1
		}
	}
	return fmt.Sprintf("%d %d", highest, lowest)
}
