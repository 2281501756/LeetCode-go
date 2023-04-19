package leetcode

import (
    "testing"
)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {

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
func deleteDuplicates(head *ListNode) *ListNode {
    h := head
    for h != nil {
        if h.Next != nil && h.Val == h.Next.Val {
            t := h.Next
            for t != nil && t.Val == h.Val {
                t = t.Next
            }
            h.Next = t
            h = h.Next
        } else {
            h = h.Next
        }
    }
    return head
}

//leetcode submit region end(Prohibit modification and deletion)
