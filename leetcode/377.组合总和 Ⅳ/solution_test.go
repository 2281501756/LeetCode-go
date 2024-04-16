package leetcode

import (
	"testing"
)

func TestCombinationSumIv(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i-nums[j] >= 0 {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

//leetcode submit region end(Prohibit modification and deletion)
