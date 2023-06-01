package leetcode

import (
	"testing"
)

func TestNThTribonacciNumber(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func tribonacci(n int) int {
	t1, t2, t3 := 0, 1, 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		for i := 3; i <= n; i++ {
			t := t1 + t2 + t3
			t1 = t2
			t2 = t3
			t3 = t
		}
		return t3
	}
}

//leetcode submit region end(Prohibit modification and deletion)
