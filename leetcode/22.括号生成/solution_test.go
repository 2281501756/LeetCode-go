package leetcode

import (
	"testing"
)

func TestGenerateParentheses(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var dfs func(l, r int, path string)
	dfs = func(l, r int, path string) {
		if l == n && r == n {
			res = append(res, path)
			return
		}
		if r > l || l > n || r > n {
			return
		}
		dfs(l+1, r, path+"(")
		dfs(l, r+1, path+")")
	}
	dfs(0, 0, "")
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
