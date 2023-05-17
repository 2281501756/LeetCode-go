package leetcode

import (
	"testing"
)

func TestSpiralMatrixIi(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	x, y, t := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		res[x][y] = i
		if x+dx[t] >= n || x+dx[t] < 0 || y+dy[t] >= n || y+dy[t] < 0 || res[x+dx[t]][y+dy[t]] != 0 {
			t = (t + 1) % 4
		}
		x += dx[t]
		y += dy[t]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
