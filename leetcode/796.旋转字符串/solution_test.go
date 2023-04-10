package leetcode

import(
    "strings"
    "testing"
)

func TestRotateString(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func rotateString(s string, goal string) bool {
    return len(s) == len(goal) && strings.Contains(s + s, goal)
}
//leetcode submit region end(Prohibit modification and deletion)
