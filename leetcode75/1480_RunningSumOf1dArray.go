package leetcode75

func runningSum(nums []int) []int {
	sum := 0
	output := []int{}

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		output = append(output, sum)
	}

	return output
}
