package leetcode75

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	pointer := 0
	for i := 0; i < len(s); i++ {
		if pointer > len(t)-1 {
			return false
		}
		for s[i] != t[pointer] {
			pointer++
			if pointer > len(t)-1 {
				return false
			}
		}
		pointer++
	}
	return true
}
