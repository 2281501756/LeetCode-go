package leetcode

import (
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	println(lengthOfLIS([]int{10, 9, 2, 5, 3, 4}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {
	dp := make([]int, 1)
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[len(dp)-1] < nums[i] {
			dp = append(dp, nums[i])
		} else {
			l, r := 0, len(dp)-1
			for l < r {
				mid := (l + r) / 2
				if dp[mid] < nums[i] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			dp[l] = nums[i]
		}
	}
	return len(dp)
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
