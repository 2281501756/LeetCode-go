package leetcode

import (
	"testing"
)

func TestBinaryGap(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func binaryGap(n int) int {
	res := 0
	for i, last := 0, -1; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				res = max(res, i-last)
			}
			last = i
		}
        n = n  >> 1
	}
    return res
}

func max(a, b int) int {
    if a >= b {
        return a
    } else {
        return b
    }
}

//leetcode submit region end(Prohibit modification and deletion)
