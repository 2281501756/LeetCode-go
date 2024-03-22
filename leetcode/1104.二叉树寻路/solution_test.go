package leetcode

import (
	"testing"
)

func TestPathInZigzagLabelledBinaryTree(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func pathInZigZagTree(label int) []int {
	row, rowStart := 1, 1
	for rowStart*2 <= label {
		row++
		rowStart *= 2
	}
	path := make([]int, 0)
	if row%2 == 0 {
		label = getReverse(label, row)
	}

	for row > 0 {
		if row%2 == 0 {
			path = append(path, getReverse(label, row))
		} else {
			path = append(path, label)
		}
		row--
		label = label >> 1
	}
	for i, n := 0, len(path); i < n/2; i++ {
		path[i], path[n-i-1] = path[n-i-1], path[i]
	}
	return path
}

func getReverse(label, row int) int {
	return (1<<row - 1)  - (label - 1<<(row-1))
}

//leetcode submit region end(Prohibit modification and deletion)
