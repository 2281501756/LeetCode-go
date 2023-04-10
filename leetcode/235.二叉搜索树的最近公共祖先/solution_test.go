package leetcode

import (
	"testing"
)

func TestLowestCommonAncestorOfABinarySearchTree(t *testing.T) {

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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p.Val > q.Val {
        p, q = q, p
    }
    for !(p.Val <= root.Val && q.Val >= root.Val) {
        if p.Val < root.Val && q.Val < root.Val {
            root = root.Left
        }
        if p.Val > root.Val && q.Val > root.Val {
            root = root.Right
        }
    }
    return root
}

//leetcode submit region end(Prohibit modification and deletion)
