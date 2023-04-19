package leetcode

import (
	"testing"
)

func TestLinkedListCycle(t *testing.T) {

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
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	quick, slow := head.Next, head
	for quick != slow {
        if quick == nil || quick.Next == nil {
            return false
        }
		quick = quick.Next.Next
		slow = slow.Next
	}
	return  true
}

//leetcode submit region end(Prohibit modification and deletion)
