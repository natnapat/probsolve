package leetcode75

func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if val, ok := m[s[i]]; ok {
			if val != t[i] {
				return false
			}
		} else {
			for _, val := range m {
				if val == t[i] {
					return false
				}
			}
			m[s[i]] = t[i]
		}
	}
	return true
}
