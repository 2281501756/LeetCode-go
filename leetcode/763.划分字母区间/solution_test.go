package leetcode

import (
	"testing"
)

func TestPartitionLabels(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func partitionLabels(s string) []int {
	hash := map[int32]int{}
	for i, k := range s {
		hash[k] = i
	}
	res := make([]int, 0)
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		end = max(end, hash[int32(s[i])])
		if end == i {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
