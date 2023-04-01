package leetcode

import (
	"testing"
)

func TestQJnOS7(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text2)+1)
	for i := range dp {
		dp[i] = make([]int, len(text1)+1)
	}
	for i := 1; i <= len(text2); i++ {
		for j := 1; j <= len(text1); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if text1[j-1] == text2[i-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	return dp[len(text2)][len(text1)]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
