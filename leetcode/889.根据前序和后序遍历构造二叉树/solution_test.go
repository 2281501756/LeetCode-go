package leetcode

import (
	"testing"
)

func TestConstructBinaryTreeFromPreorderAndPostorderTraversal(t *testing.T) {
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	if len(preorder) == 1 {
		return root
	}
	idx := 0
	for i := 0; i < len(preorder); i++ {
		if postorder[i] != preorder[1] {
			continue
		}
		idx = i + 1
		break
	}
	root.Left = constructFromPrePost(preorder[1:idx+1], postorder[:idx])
	root.Right = constructFromPrePost(preorder[idx+1:], postorder[idx:])
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
