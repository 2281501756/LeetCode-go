package leetcode

import(
    "testing"
)

func TestM1oyTv(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
    smap, tmap := map[byte]int{}, map[byte]int{}
    ansL, ansR := -1, -1
    for i := 0; i < len(t); i++ {
        tmap[t[i]]++
    }
    check := func() bool {
        for k, v := range tmap {
            if smap[k] < v {
                return false
            }
        }
        return true
    }
    for l, r := 0, 0; r < len(s); r++ {
        smap[s[r]]++
        for check() {
            if ansL == -1 && ansR == -1 || r-l < ansR-ansL {
                ansL = l
                ansR = r
            }
            smap[s[l]]--
            l++
        }
    }
    if ansL == -1 {
        return ""
    }
    return s[ansL : ansR+1]
}
//leetcode submit region end(Prohibit modification and deletion)
