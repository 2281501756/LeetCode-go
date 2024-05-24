package leetcode

import (
	"testing"
)

func TestFindTheMostCompetitiveSubsequence(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && len(nums)-i+len(stack) > k && nums[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return stack[:k]
}

//leetcode submit region end(Prohibit modification and deletion)
