package leetcode

import (
	"testing"
)

func TestCountSortedVowelStrings(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func countVowelStrings(n int) int {
    dp := []int{1,1,1,1,1}
    for i := 1; i < n; i ++ {
        for j := 1; j < 5; j ++ {
			dp[j] += dp[j - 1]
		}
    }
	res := 0
	for _, k := range dp {
		res += k
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
