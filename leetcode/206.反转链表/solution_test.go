package leetcode

import (
	"testing"
)

func TestReverseLinkedList(t *testing.T) {

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
	var h, t *ListNode
	h, t = head, nil
	for h != nil {
        next := h.Next
        h.Next = t
        t = h
        h = next
	}
	return t
}

//leetcode submit region end(Prohibit modification and deletion)
