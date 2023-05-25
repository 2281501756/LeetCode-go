package leetcode

import (
	"testing"
)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    t := head
    leng := 0
    for t != nil {
        t = t.Next
        leng++
    }
    t = head
    if leng - n == 0 {
        return head.Next
    }
    for i := 0; i < leng - n - 1 ; i ++ {
        t = t.Next
    }
    t.Next = t.Next.Next
    return head
}

//leetcode submit region end(Prohibit modification and deletion)
