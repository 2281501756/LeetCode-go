package leetcode

import (
    "testing"
)

func TestSortList(t *testing.T) {

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
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	temp := slow.Next
	slow.Next = nil
	left, right := sortList(head), sortList(temp)
	h := &ListNode{0, nil}
	t := h
	for left != nil && right != nil {
		if left.Val < right.Val {
			t.Next = left
			t = t.Next
			left = left.Next
		} else {
			t.Next = right
			t = t.Next
			right = right.Next
		}
	}
	for left != nil {
		t.Next = left
		t = t.Next
		left = left.Next
	}
	for right != nil {
		t.Next = right
		t = t.Next
		right = right.Next
	}
	return h.Next
}

//leetcode submit region end(Prohibit modification and deletion)
