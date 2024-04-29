package leetcode

import (
	"fmt"
	"testing"
)

func TestConvertToBase2(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func baseNeg2(n int) string {
	i, j := 1, 0

	for i <= n {
		if j&1 == 1 && i&n > 0 {
			n += i << 1
		}
		i <<= 1
		j++
	}
	return fmt.Sprintf("%0b", n)
}

//leetcode submit region end(Prohibit modification and deletion)
