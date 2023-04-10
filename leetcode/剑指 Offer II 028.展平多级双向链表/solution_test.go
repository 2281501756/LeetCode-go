package leetcode

import (
	"testing"
)

func TestQv1Da2(t *testing.T) {

}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	dfs(root)
	return root
}
func dfs(root *Node) (last *Node) {
	cur := root
	for cur != nil {
		next := cur.Next
		if cur.Child != nil {
			lastchild := dfs(cur.Child)
			cur.Next = cur.Child
			cur.Child.Prev = cur
			if next != nil {
				lastchild.Next = next
				next.Prev = lastchild
			}
			cur.Child = nil
			last = lastchild
		} else {
			last = cur
		}
		cur = next
	}
	return last
}

//leetcode submit region end(Prohibit modification and deletion)
