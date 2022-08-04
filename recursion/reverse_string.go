package recursion

func Reverse_string(input string) string {
	//fmt.Println(input)
	if input == "" {
		return ""
	}
	return Reverse_string(string(input[1:])) + string(input[0])
}
