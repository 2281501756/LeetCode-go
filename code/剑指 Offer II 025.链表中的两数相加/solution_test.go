package leetcode

import (
	"testing"
)

func TestLMSNwu(t *testing.T) {

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
	s1, s2 := make([]int, 0), make([]int, 0)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	t := 0
	var res *ListNode
	for len(s1) != 0 || len(s2) != 0 || t != 0 {
		top1, top2 := 0, 0
		if len(s1) != 0 {
			top1 = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) != 0 {
			top2 = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		t = t + top1 + top2
		node := &ListNode{
			Val:  t % 10,
			Next: res,
		}
		t = t / 10
		res = node
	}
	return res

}

//leetcode submit region end(Prohibit modification and deletion)
