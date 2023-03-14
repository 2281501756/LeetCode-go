package leetcode

import (
	"testing"
)

func TestC32eOV(t *testing.T) {

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
func detectCycle(head *ListNode) *ListNode {
	mmap := map[*ListNode]bool{}
	for head != nil {
		if _, has := mmap[head]; has {
			return head
		}
		mmap[head] = true
		head = head.Next
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
