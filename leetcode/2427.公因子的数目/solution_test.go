package leetcode

import(
    "testing"
)

func TestNumberOfCommonFactors(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func commonFactors(a int, b int) int {
    x := min(a, b)
    res := 0
    for i := 1; i <= x; i ++ {
        if a % i == 0 && b % i == 0 {
            res ++
        }
    }
    return res
}
func min (a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}
//leetcode submit region end(Prohibit modification and deletion)
