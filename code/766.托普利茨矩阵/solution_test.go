package leetcode

import (
	"testing"
)

func TestToeplitzMatrix(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func isToeplitzMatrix(matrix [][]int) bool {
    n, m := len(matrix), len(matrix[0])
    for i := 0; i < m; i ++ {
        x, y := 0, i
        prev := -1
        for x < n && y < m {
            if prev == -1 {
                prev = matrix[x][y]
            } else {
                if prev != matrix[x][y] {
                    return false
                }
            }
            x = x + 1
            y = y + 1
        }
    }
    for i := 0; i < n; i ++ {
        x, y := i, 0
        prev := -1
        for x < n && y < m {
            if prev == -1 {
                prev = matrix[x][y]
            } else {
                if prev != matrix[x][y] {
                    return false
                }
            }
            x = x + 1
            y = y + 1
        }
    }
    return true
}

//leetcode submit region end(Prohibit modification and deletion)
