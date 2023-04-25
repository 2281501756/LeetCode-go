package leetcode

import (
	"testing"
)

func TestLeafSimilarTrees(t *testing.T) {

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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(node *TreeNode)
	arr, arr1, arr2 := make([]int, 0), make([]int, 0), make([]int, 0)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			arr = append(arr, root.Val)
		}

		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root1)
	arr1 = append([]int{}, arr...)
	arr = []int{}
	dfs(root2)
	arr2 = append([]int{}, arr...)
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
