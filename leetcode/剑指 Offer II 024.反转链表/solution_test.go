package leetcode

import (
	"testing"
)

func TestUHnkqh(t *testing.T) {

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
func reverseList(head *ListNode) *ListNode {
	var prev, h1 *ListNode
	prev, h1 = nil, head
	for h1 != nil {
		t := h1.Next
		h1.Next = prev
		prev = h1
		h1 = t
	}
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)
