package leetcode

import (
	"testing"
)

func TestMergeKSortedLists(t *testing.T) {

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
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists) - 1)
}
func merge(list []*ListNode, l, r int) *ListNode {
	if l > r {
		return nil
	}
	if l == r {
		return list[l]
	}
	mid := (l + r) >> 1
	return mergeTwoList(merge(list, l, mid), merge(list, mid+1, r))
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
			tail = tail.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
			tail = tail.Next
		}
	}
	if l1 != nil {
		tail.Next = l1
		l1 = l1.Next
		tail = tail.Next
	}
	if l2 != nil {
		tail.Next = l2
		l2 = l2.Next
		tail = tail.Next
	}
	return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
