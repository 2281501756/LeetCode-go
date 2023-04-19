package leetcode

import (
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	t := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			t.Next = list1
			list1 = list1.Next
			t = t.Next
		} else{
            t.Next = list2
            list2 = list2.Next
            t = t.Next
        }
	}
    for list1 != nil {
        t.Next = list1
        list1 = list1.Next
        t = t.Next
    }
    for list2 != nil {
        t.Next = list2
        list2 = list2.Next
        t = t.Next
    }
    return head.Next
}

//leetcode submit region end(Prohibit modification and deletion)
