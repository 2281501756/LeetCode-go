package leetcode

import (
	"sort"
	"testing"
)

func TestSortTheMatrixDiagonally(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func diagonalSort(mat [][]int) [][]int {

	n, m := len(mat), len(mat[0])
	for i := 0; i < n+m-1; i++ {
		x, y := 0, 0
		if i < m {
			x = 0
			y = m - i - 1
		} else {
			x = i - m + 1
			y = 0
		}

		arr := make([]int, 0)
		for x < n && y < m {
			arr = append(arr, mat[x][y])
			x++
			y++
		}
		sort.Ints(arr)

		if i < m {
			x = 0
			y = m - i - 1
		} else {
			x = i - m + 1
			y = 0
		}
		j := 0
		for x < n && y < m {
			mat[x][y] = arr[j]
			x++
			y++
			j++
		}
	}
	return mat
}

//leetcode submit region end(Prohibit modification and deletion)
