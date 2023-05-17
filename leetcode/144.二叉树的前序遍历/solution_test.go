package leetcode

import (
	"testing"
)

func TestBinaryTreePreorderTraversal(t *testing.T) {

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
func preorderTraversal(root *TreeNode) []int {
	var dfs func(node *TreeNode)
	res := make([]int, 0)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		res = append(res, r.Val)
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
