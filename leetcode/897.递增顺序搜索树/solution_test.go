package leetcode

import (
	"testing"
)

func TestIncreasingOrderSearchTree(t *testing.T) {

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
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	list := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		list = append(list, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	res := &TreeNode{}
	t := res
	for i := 0; i < len(list); i++ {
		t.Right = &TreeNode{Val: list[i]}
		t = t.Right
	}
	return res.Right
}

//leetcode submit region end(Prohibit modification and deletion)
