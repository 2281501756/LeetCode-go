package leetcode

import (
	"sort"
	"testing"
)

func TestCoinChangeIi(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func change(amount int, coins []int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})

	dp := make([][]int, amount+1)
	for i := range dp {
		dp[i] = make([]int, len(coins))
	}
	for i := 0; i < len(coins); i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			dp[i][j] = 0
			for z := j; z >= 0; z-- {
				if i-coins[z] < 0 {
					continue
				}
				dp[i][j] += dp[i-coins[z]][z]
			}
		}
	}
	return dp[amount][len(coins)-1]
}

//leetcode submit region end(Prohibit modification and deletion)
