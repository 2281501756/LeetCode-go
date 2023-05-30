package leetcode

import (
	"testing"
)

func TestDeleteNodesAndReturnForest(t *testing.T) {

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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	s := make([]bool, 1010)
	for _, i := range to_delete {
		s[i] = true
	}
	res := make([]*TreeNode, 0)
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		root.Left = dfs(root.Left)
		root.Right = dfs(root.Right)
		isDelete := s[root.Val]
		if !isDelete {
			return root
		}
		if root.Left != nil {
			res = append(res, root.Left)
		}
		if root.Right != nil {
			res = append(res, root.Right)
		}
        return nil
	}
    t := dfs(root)
    if t != nil {
        res =  append(res, t)
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
