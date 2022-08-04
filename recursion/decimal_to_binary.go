package recursion

import (
	"strconv"
)

func Decimal_to_binary(decimal int, result string) string {
	if decimal == 0 {
		return result
	}
	result = strconv.Itoa(decimal%2) + result
	return Decimal_to_binary(decimal/2, result)
}
