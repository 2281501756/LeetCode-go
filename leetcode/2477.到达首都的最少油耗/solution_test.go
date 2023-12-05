package leetcode

import (
	"testing"
)

func TestMinimumFuelCostToReportToTheCapital(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumFuelCost(roads [][]int, seats int) int64 {
	g := make([][]int, len(roads)+1)
	for _, e := range roads {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	var res int64
	dfs = func(root int, except int) int {
		node := g[root]
		s := 1
		for i := 0; i < len(node); i++ {
			if node[i] != except {
				s += dfs(node[i], root)
			}
		}
		if root != 0 {
			res += int64((s-1)/seats + 1)
		}
		return s
	}
	dfs(0, -1)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
