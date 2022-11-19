package line

//run-length encoded
type RleLines []int

func(l RleLines) Get(i int) int {
	j := 0
	for i > 0 {
		i -= l[j]

		if i < 0 {
			return l[j + 1]
		}

		j += 2
	}

	return l[j + 1]
}
