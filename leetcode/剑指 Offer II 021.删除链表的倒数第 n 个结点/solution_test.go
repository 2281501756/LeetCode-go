package leetcode

import (
	"testing"
)

func TestSLwz0R(t *testing.T) {

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
	if head == nil {
		return nil
	}
	length := getLength(head)
	t := length - n
	tail := head
	if t == 0 {
		return head.Next
	}
	for i := 0; i < t-1; i++ {
		tail = tail.Next
	}
	tail.Next = tail.Next.Next
	return head
}
func getLength(root *ListNode) (length int) {
	for root != nil {
		length++
		root = root.Next
	}
	return length
}

//leetcode submit region end(Prohibit modification and deletion)
