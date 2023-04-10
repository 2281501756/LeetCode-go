package leetcode

import (
	"testing"
)

func TestErWeiShuZuZhongDeChaZhaoLcof(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    x,y := 0, len(matrix[0]) - 1
    for x < len(matrix) && y >= 0 {
        if matrix[x][y] == target {
            return true
        } else if matrix[x][y] > target {
            y--
        } else {
            x++
        }
    }
    return false
}

//leetcode submit region end(Prohibit modification and deletion)
