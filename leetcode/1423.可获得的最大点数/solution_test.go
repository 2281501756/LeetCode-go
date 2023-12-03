package leetcode

import (
	"testing"
)

func TestMaximumPointsYouCanObtainFromCards(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxScore(cardPoints []int, k int) int {
	n, sum, cur := len(cardPoints), 0, 0
	for i := 0; i < n; i++ {
		sum += cardPoints[i]
	}
	for i := 0; i < n-k; i++ {
		cur += cardPoints[i]
	}
	res := cur
	for i := n - k; i < n; i++ {
		cur += cardPoints[i]
		cur -= cardPoints[i-n+k]
		res = min(res, cur)
	}
	return sum - res
}

//leetcode submit region end(Prohibit modification and deletion)
