package calc

func GenerateInteger() func() int {
	i := 0

	return func() int {
		i += 1
		return i
	}

}
