package leetcode

import(
    "testing"
)

func TestDKk3P7(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
    if s == t || len(s) != len(t) {
        return false
    }
    mmap := make(map[int32]int, 0)
    for _, v := range s {
        mmap[v]++
    }
    for _, v := range t {
        mmap[v]--
        if mmap[v] == 0 {
            delete(mmap, v)
        }
    }
    return len(mmap) == 0
}
//leetcode submit region end(Prohibit modification and deletion)
