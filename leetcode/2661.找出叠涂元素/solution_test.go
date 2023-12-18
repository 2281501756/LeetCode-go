package leetcode

import (
	"testing"
)

func TestFirstCompletelyPaintedRowOrColumn(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func firstCompleteIndex(arr []int, mat [][]int) int {
	n, m := len(mat), len(mat[0])
	mp := map[int][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mp[mat[i][j]] = []int{i, j}
		}
	}
	row, col := make([]int, n), make([]int, m)
	for i := 0; i < len(arr); i++ {
		address := mp[arr[i]]
		x, y := address[0], address[1]
		row[x]++
		col[y]++
		if row[x] >= m || col[y] >= n {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
