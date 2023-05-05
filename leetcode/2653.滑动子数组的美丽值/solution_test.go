package leetcode

import (
	"fmt"
	"testing"
)

func TestSlidingSubarrayBeauty(t *testing.T) {
	fmt.Println(getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2))
}

// leetcode submit region begin(Prohibit modification and deletion)
func getSubarrayBeauty(nums []int, k int, x int) []int {
	res := make([]int, len(nums)-k+1)
	q := make([]int, 110)
	for _, num := range nums[:k-1] {
		q[num+50]++
	}
	for i, num := range nums[k-1:] {
		q[num+50]++
		left := x
		for j, c := range q[:50] {
			left -= c
			if left <= 0 {
				res[i] = j - 50
				break
			}
		}
		q[nums[i]+50]--
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
