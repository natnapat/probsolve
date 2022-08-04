package recursion

func Binary_search(arr []int, left int, right int, target int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if arr[mid] == target {
		return mid
	}

	if target < arr[mid] {
		return Binary_search(arr, left, mid-1, target)
	}
	return Binary_search(arr, mid+1, right, target)
}
