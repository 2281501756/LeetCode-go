package leetcode

import (
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-2], dp[i-1])
		if i < len(cost) {
			dp[i] += cost[i]
		}
	}
	return dp[len(cost)]
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
