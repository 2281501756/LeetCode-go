package leetcode

import (
	"testing"
)

func TestNextGreaterNodeInLinkedList(t *testing.T) {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	st, tt := make([]int, len(nums)), 0
	for i :=len(nums) - 1; i >= 0; i -- {
		for tt > 0 && st[tt - 1] <= nums[i] {
			tt --
		}
		t := nums[i]
		if tt == 0 {
			nums[i] = 0
		} else {
			nums[i] = st[tt - 1]
		}
		st[tt] = t
		tt++
	}
	return nums

}

//leetcode submit region end(Prohibit modification and deletion)
