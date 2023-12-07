package leetcode

import (
	"testing"
)

func TestReorderRoutesToMakeAllPathsLeadToTheCityZero(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minReorder(n int, connections [][]int) int {
	g := make([][][2]int, n)
	for _, item := range connections {
		x1, x2 := item[0], item[1]
		g[x1] = append(g[x1], [2]int{x2, 1})
		g[x2] = append(g[x2], [2]int{x1, 0})
	}
	var dfs func(int, int) int
	dfs = func(root int, except int) int {
		res := 0
		for _, item := range g[root] {
			node, flag := item[0], item[1]
			if node != except {
				res += dfs(node, root)
				res += flag
			}
		}
		return res
	}
	return dfs(0, -1)
}

//leetcode submit region end(Prohibit modification and deletion)
