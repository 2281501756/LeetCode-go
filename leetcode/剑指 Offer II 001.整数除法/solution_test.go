package leetcode

import (
    "fmt"
    "math"
	"testing"
)

func TestXoh6Oh(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func divide(a int, b int) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	sign := 1
	if a < 0 && b > 0 || a > 0 && b < 0 {
		sign = -1
	}
	res := 0
	if a > 0 {
		a = -a
	}
	if b > 0 {
		b = -b
	}
	for a <= b {
		value, k := b, 1
		for a-value < value {
			value = value << 1
			k = k << 1
		}
		a -= value
		res += k

	}
	if sign == -1 {
		return -res
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
