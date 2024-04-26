package leetcode

import (
	"testing"
)

func TestClimbingStairs(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
