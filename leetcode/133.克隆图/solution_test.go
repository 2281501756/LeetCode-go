package leetcode

import (
	"testing"
)

func TestCloneGraph(t *testing.T) {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	nodeMap := make(map[int]*Node)
	queue := []*Node{node}

	nodeMap[node.Val] = &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		copyNode := nodeMap[head.Val]
		for _, neighbor := range head.Neighbors {
			if _, ok := nodeMap[neighbor.Val]; !ok {
				queue = append(queue, neighbor)
				neighborCopy := &Node{Val: neighbor.Val, Neighbors: make([]*Node, 0)}
				nodeMap[neighbor.Val] = neighborCopy
			}
			copyNode.Neighbors = append(copyNode.Neighbors, nodeMap[neighbor.Val])
		}
	}
	return nodeMap[node.Val]
}

//leetcode submit region end(Prohibit modification and deletion)
