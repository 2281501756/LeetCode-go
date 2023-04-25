package leetcode

import (
	"testing"
)

func TestProjectionAreaOf3dShapes(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func projectionArea(grid [][]int) int {
	res := 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != 0 {
				res++
			}
		}
	}
	for i := 0; i < n; i++ {
		t := 0
		for j := 0; j < m; j++ {
			if grid[i][j] > t {
				t = grid[i][j]
			}
		}
		res += t
	}
	for j := 0; j < m; j++ {
		t := 0
		for i := 0; i < n; i++ {
			if grid[i][j] > t {
				t = grid[i][j]
			}
		}
		res += t
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
