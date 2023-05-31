package leetcode

import (
	"testing"
)

func TestFeiBoNaQiShuLieLcof(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		t := (a + b) % (1e9 + 7)
		a = b
		b = t
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
