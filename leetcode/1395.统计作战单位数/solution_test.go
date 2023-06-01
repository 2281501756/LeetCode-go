package leetcode

import (
	"testing"
)

func TestCountNumberOfTeams(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numTeams(rating []int) int {
	res := 0
	for j := 1; j < len(rating)-1; j++ {
		lless, lmore, rless, rmore := 0, 0, 0, 0
		for l := 0; l < j; l++ {
			if rating[l] < rating[j] {
				lless++
			} else if rating[l] > rating[j] {
				lmore++
			}
		}
		for r := j + 1; r < len(rating); r++ {
			if rating[r] < rating[j] {
				rless++
			} else if rating[r] > rating[j] {
				rmore++
			}
		}
		res += lless*rmore + lmore*rless
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
