package leetcode

import (
	"math"
	"testing"
)

func TestBalancedBinaryTree(t *testing.T) {

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
func isBalanced(root *TreeNode) bool {
	res := true
	var dfs func(node *TreeNode) float64
	dfs = func(r *TreeNode) float64 {
		if r == nil {
			return 0
		}
		lh, rh := dfs(r.Left), dfs(r.Right)
		if math.Abs(lh-rh) > 1 {
			res = false
		}
		return max(lh, rh) + 1
	}
    dfs(root)
    return res
}
func max(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
