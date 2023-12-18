package leetcode

import (
	"testing"
)

func TestCarPooling(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func carPooling(trips [][]int, capacity int) bool {
	arr := make([]int, 1010)
	for _, trip := range trips {
		arr[trip[1]] += trip[0]
		arr[trip[2]] -= trip[0]
	}
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i] + arr[i-1]
		res = max(res, arr[i])
	}
	return res <= capacity
}

//leetcode submit region end(Prohibit modification and deletion)
