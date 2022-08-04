package recursion

func Sum_of_natural_numbers(input int) int {
	if input <= 0 {
		return input
	}
	return input + Sum_of_natural_numbers(input-1)
}
