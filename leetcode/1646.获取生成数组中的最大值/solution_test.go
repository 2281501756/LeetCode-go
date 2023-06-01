package leetcode

import (
	"testing"
)

func TestGetMaximumInGeneratedArray(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[1] = 1
	res := 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i/2] + dp[i/2+1]
			if dp[i] > res {
				res = dp[i]
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
