package leetcode

import (
	"testing"
)

func TestDistributeCoinsInBinaryTree(t *testing.T) {

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
func distributeCoins(root *TreeNode) (res int) {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		s := dfs(root.Left) + dfs(root.Right) + root.Val - 1
		res += abs(s)
		return s
	}
	dfs(root)
	return
}
func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}

//leetcode submit region end(Prohibit modification and deletion)
