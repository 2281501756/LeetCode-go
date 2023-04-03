package leetcode

import(
    "strings"
    "testing"
)

func TestJewelsAndStones(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func numJewelsInStones(jewels string, stones string) int {
    res := 0
    for i := range stones {
        if strings.Contains(jewels, string(stones[i])) {
            res ++
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
