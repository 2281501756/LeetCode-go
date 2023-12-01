package leetcode

import (
	"testing"
)

func TestAlternatingDigitSum(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func alternateDigitSum(n int) int {
	arr := []int{}
	for ; n > 0; n = n / 10 {
		arr = append(arr, n%10)
	}
	res := 0
	for i, symbol := len(arr)-1, 1; i >= 0; i, symbol = i-1, -symbol {
		res += arr[i] * symbol
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
