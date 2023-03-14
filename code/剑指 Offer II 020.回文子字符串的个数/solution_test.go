package leetcode

import (
	"testing"
)

func TestA7VOhD(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j, k := i, i; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] == s[k] {
				res++
			} else {
				break
			}
		}
		for j, k := i, i+1; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] == s[k] {
				res++
			} else {
				break
			}
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
