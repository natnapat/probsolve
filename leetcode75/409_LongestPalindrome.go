package leetcode75

func longestPalindrome(s string) int {
	hmap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := hmap[s[i]]; ok {
			hmap[s[i]]++
		} else {
			hmap[s[i]] = 1
		}
	}
	result := 0
	odd := false
	for _, v := range hmap {
		if odd == false && v%2 == 1 {
			result++
			odd = true
		}
		result += (v / 2) * 2
	}
	return result
}
