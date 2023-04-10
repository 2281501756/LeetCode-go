package leetcode

import (
	"testing"
)

func TestDiagonalTraverse(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findDiagonalOrder(mat [][]int) (ans []int) {
	for m, n, k := len(mat), len(mat[0]), 0; k < m+n-1; k++ {
		if k&1 == 1 {
			for x := max(0, k-n+1); x < min(k+1, m); x++ {
				ans = append(ans, mat[x][k-x])
			}
		} else {
			for x := min(k, m-1); x >= max(0, k-n+1); x-- {
				ans = append(ans, mat[x][k-x])
			}
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
