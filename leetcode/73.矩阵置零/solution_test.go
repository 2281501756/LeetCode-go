package leetcode

import (
    "fmt"
    "testing"
)

func TestSetMatrixZeroes(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
    rowHas0, colHas0 := false, false
    n, m := len(matrix), len(matrix[0])
    for i := 0; i < n; i++ {
        if matrix[i][0] == 0 {
            rowHas0 = true
            break
        }
    }
    for i := 0; i < m; i++ {
        if matrix[0][i] == 0 {
            colHas0 = true
            break
        }
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][j] == 0 {
                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }
    for i := 1; i < n; i ++ {
        if matrix[i][0] == 0 {
            setRowZero(matrix, i)
        }
    }
    for i := 1; i < m; i ++ {
        if matrix[0][i] == 0 {
            setColumnZero(matrix, i)
        }
    }
    fmt.Print(rowHas0, colHas0)
    if rowHas0 {
       setColumnZero(matrix, 0)
    }
    if colHas0 {
       setRowZero(matrix, 0)
    }

}
func setRowZero(arr [][]int, n int) {
    for i := 0; i < len(arr[0]); i ++ {
        arr[n][i] = 0
    }
}
func setColumnZero(arr [][]int, n int) {
    for i := 0; i < len(arr); i ++ {
        arr[i][n] = 0
    }
}
//leetcode submit region end(Prohibit modification and deletion)
