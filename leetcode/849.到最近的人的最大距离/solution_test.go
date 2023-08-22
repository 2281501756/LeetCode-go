package leetcode

import (
	"testing"
)

func TestMaximizeDistanceToClosestPerson(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxDistToClosest(seats []int) int {
	arr := make([]int, 0)
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			arr = append(arr, i)
		}
	}
	res := 0
	if len(arr) == 1 {
		return max(arr[0], len(seats)-arr[0]-1)
	}
	for i := 1; i < len(arr); i++ {
		res = max(res, (arr[i]-arr[i-1])/2)
	}
	if arr[0] != 0 {
		res = max(res, arr[0])
	}
	if arr[len(arr)-1] != len(seats)-1 {
		res = max(res, len(seats)-arr[len(arr)-1]-1)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
