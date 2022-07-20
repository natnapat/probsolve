package leetcode75

func pivotIndex(nums []int) int {
	sum := 0
	leftsum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if leftsum == sum-nums[i]-leftsum {
			return i
		}

		leftsum += nums[i]
	}
	return -1

}
