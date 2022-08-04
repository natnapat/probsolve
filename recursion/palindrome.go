package recursion

func IsPalindrome(input string) bool {
	//define base-case / stopping condition
	if len(input) == 0 || len(input) == 1 {
		return true
	}

	//shrink the problem space
	if input[0] == input[len(input)-1] {
		return IsPalindrome(string(input[1 : len(input)-1]))
	}

	//additional base-case to handle non-palindrome
	return false
}
