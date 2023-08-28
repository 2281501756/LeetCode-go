package leetcode

import (
	"testing"
)

func TestInsertInterval(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	isMerged := false
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			res = append(res, interval)
		} else if interval[0] > newInterval[1] {
			if !isMerged {
				isMerged = true
				res = append(res, newInterval)
			}
			res = append(res, interval)
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	if !isMerged {
		res = append(res, newInterval)
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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
