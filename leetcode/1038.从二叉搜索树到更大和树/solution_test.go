package leetcode

import (
	"testing"
)

func TestBinarySearchTreeToGreaterSumTree(t *testing.T) {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	s := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		s += node.Val
		node.Val = s
		dfs(node.Left)
	}
	dfs(root)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
