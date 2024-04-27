package leetcode

import (
	"testing"
)

func TestFindTheWidthOfColumnsOfAGrid(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findColumnWidth(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	res := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[i] = max(res[i], getLen(grid[j][i]))
		}
	}
	return res
}

func getLen(n int) (res int) {
	if n == 0 {
		return 1
	}
	if n < 0 {
		res++
		n = -n
	}
	for n > 0 {
		n = n / 10
		res++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
