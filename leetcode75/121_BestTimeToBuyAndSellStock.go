package leetcode75

// func maxProfit(prices []int) int {
//     maxp := 0
//     for i:= 0; i<len(prices); i++ {
//         for j:= i+1; j<len(prices); j++ {
//             if prices[j]-prices[i] > maxp {
//                 maxp = prices[j]-prices[i]
//             }
//         }
//     }
//     return maxp
// }

func maxProfit(prices []int) int {
	maxp := 0
	left := 0
	right := 1

	for right < len(prices) {
		if prices[left] < prices[right] {
			if prices[right]-prices[left] > maxp {
				maxp = prices[right] - prices[left]
			}
		} else {
			left = right
		}
		right++
	}
	return maxp
}
