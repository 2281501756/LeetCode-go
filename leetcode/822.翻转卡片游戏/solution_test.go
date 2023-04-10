package leetcode

import (
	"testing"
)

func TestCardFlippingGame(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func flipgame(fronts []int, backs []int) int {
	mmap := make(map[int]bool, 0)
	for i, num := range fronts {
		if num == backs[i] {
            mmap[num] = true
		}
	}
    res := 2001
    for i := range fronts {
        if mmap[fronts[i]] == false {
            res = min(res, fronts[i])
        }
        if mmap[backs[i]] == false {
            res = min(res, backs[i])
        }
    }
    if res == 2001 {
        res = 0
    }
    return res
}
func min(a, b int) int  {
    if a < b {
        return a
    } else {
        return b
    }
}

//leetcode submit region end(Prohibit modification and deletion)
