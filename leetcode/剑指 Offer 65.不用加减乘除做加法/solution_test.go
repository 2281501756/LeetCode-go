package leetcode

import(
    "testing"
)

func TestBuYongJiaJianChengChuZuoJiaFaLcof(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func add(a int, b int) int {
    if b == 0 {
        return a
    }
    return add(a ^ b, (a & b) << 1)
}
//leetcode submit region end(Prohibit modification and deletion)
