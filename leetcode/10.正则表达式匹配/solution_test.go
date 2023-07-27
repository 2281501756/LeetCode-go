package leetcode

import (
	"testing"
)

func TestRegularExpressionMatching(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	s = " " + s
	p = " " + p
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i > 0 && p[j] != '*' {
				f[i][j] = f[i-1][j-1] && (s[i] == p[j] || p[j] == '.')
			} else if p[j] == '*' {
				f[i][j] = f[i][j-2] || i > 0 && f[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return f[n][m]
}

//leetcode submit region end(Prohibit modification and deletion)
