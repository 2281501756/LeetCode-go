package leetcode

import (
	"testing"
)

func TestInsufficientNodesInRootToLeafPaths(t *testing.T) {

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
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	limit -= root.Val
	if root.Left == nil && root.Right == nil {
		if limit > 0 {
			return nil
		} else {
			return root
		}
	}
	root.Left = sufficientSubset(root.Left, limit)
	root.Right = sufficientSubset(root.Right, limit)
    if root.Left == nil && root.Right == nil {
        return nil
    } else {
        return root
    }
}

//leetcode submit region end(Prohibit modification and deletion)
