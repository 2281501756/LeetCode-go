package leetcode

import (
	"testing"
)

func TestLGjMqU(t *testing.T) {

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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	list := make([]*ListNode, 0)
	for h := head; h != nil; h = h.Next {
		list = append(list, h)
	}
	i, j := 0, len(list)-1
	for i < j {
		list[i].Next = list[j]
		i++
		if i == j {
			break
		}
		list[j].Next = list[i]
		j--
	}
	list[i].Next = nil
}

//leetcode submit region end(Prohibit modification and deletion)
