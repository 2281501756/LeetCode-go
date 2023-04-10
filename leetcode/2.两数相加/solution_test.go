package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{0, nil}
	current := res
	t := 0
	for l1 != nil || l2 != nil || t != 0{
		num1, num2 := 0, 0
		if l1 != nil {
			num1 = l1.Val
            l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
            l2 = l2.Next
		}
        t = num1 + num2 + t
		current.Next = &ListNode{t % 10, nil}
        current = current.Next
        t = t / 10
	}
    return res.Next
}

//leetcode submit region end(Prohibit modification and deletion)
