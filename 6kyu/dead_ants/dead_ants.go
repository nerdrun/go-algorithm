package dead_ants

func DeadAntCount(ants string) int {
	a := 0
	n := 0
	t := 0

	for i := 0; i < len(ants); i++ {
		char := ants[i]
		if char == '.' {
			continue
		}
		if i < len(ants)-2 && ants[i:i+3] == "ant" {
			i += 2
			continue
		}

		if char == 'a' {
			a++
		} else if char == 'n' {
			n++
		} else if char == 't' {
			t++
		}
	}

	deathCount := a
	if a < n {
		deathCount = n
	}
	if a < t {
		deathCount = t
	}

	return deathCount
}
