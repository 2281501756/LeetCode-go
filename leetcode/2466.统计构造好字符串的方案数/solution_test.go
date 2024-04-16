package leetcode

import (
	"testing"
)

func TestCountWaysToBuildGoodStrings(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	for i := 0; i <= high; i++ {
		if i-zero >= 0 {
			dp[i] += dp[i-zero]
		}
		dp[i] %= (1e9+7)
		if i-one >= 0 {
			dp[i] += dp[i-one]
		}
		dp[i] %= (1e9+7)
	}
	res := 0
	for i := low; i <= high; i++ {
		res += dp[i]
	}
	return res % (1e9+7)
}

//leetcode submit region end(Prohibit modification and deletion)
