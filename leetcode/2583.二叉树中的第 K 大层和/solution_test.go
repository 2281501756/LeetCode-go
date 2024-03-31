package leetcode

import (
	"sort"
	"testing"
)

func TestKthLargestSumInABinaryTree(t *testing.T) {

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
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}
	levelSum := make([]int, 0)
	q := []*TreeNode{root}
	l := 0
	for l < len(q) {
		sum := 0
		oldR := len(q)
		for l < oldR {
			sum += q[l].Val
			if q[l].Left != nil {
				q = append(q, q[l].Left)
			}
			if q[l].Right != nil {
				q = append(q, q[l].Right)
			}
			l++
		}
		levelSum = append(levelSum, sum)
	}
	sort.Slice(levelSum, func(n, j int) bool {
		return levelSum[n] > levelSum[j]
	})

	if k-1 >= len(levelSum) {
		return -1
	} else {
		return int64(levelSum[k-1])
	}
}

//leetcode submit region end(Prohibit modification and deletion)
