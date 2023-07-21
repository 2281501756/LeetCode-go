package leetcode

import (
	"testing"
)

func TestMaximumSumCircularSubarray(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubarraySumCircular(nums []int) int {
	total, maxSum, minSum, currMax, currmin := nums[0], nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		total += nums[i]
		currMax = max(nums[i], currMax+nums[i])
		maxSum = max(maxSum, currMax)
		currmin = min(nums[i], currmin+nums[i])
		minSum = min(minSum, currmin)
	}
	if total == minSum {
		return maxSum
	} else {
		return max(maxSum, total-minSum)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
