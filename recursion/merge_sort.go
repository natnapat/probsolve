package recursion

func Merge_sort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := Merge_sort(items[:len(items)/2])
	second := Merge_sort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	result := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}
	return result
}
