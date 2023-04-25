package leetcode

import(
    "testing"
)

func TestBinaryTreePruning(t *testing.T){
    
}
 type TreeNode struct {
         Val int
         Left *TreeNode
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
func pruneTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left = pruneTree(root.Left)
    root.Right = pruneTree(root.Right)
    if root.Val == 0 && root.Left == nil && root.Right == nil {
        return nil
    }
    return root
}
//leetcode submit region end(Prohibit modification and deletion)
