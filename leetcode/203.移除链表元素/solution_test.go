package leetcode

import (
	"testing"
)

func TestRemoveLinkedListElements(t *testing.T) {

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
func removeElements(head *ListNode, val int) *ListNode {
    h := &ListNode{0, head}
    t := h
    for t != nil {
        if t.Next != nil && t.Next.Val == val {
            t.Next = t.Next.Next
        } else {
            t = t.Next
        }
    }
    return h.Next
}

//leetcode submit region end(Prohibit modification and deletion)
