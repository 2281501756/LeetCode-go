package leetcode

import (
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(res) {
			res = s[l+1 : r]
		}
	}
    for i := 0; i < len(s); i ++ {
        l, r := i, i + 1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            l--
            r++
        }
        if r-l-1 > len(res) {
            res = s[l+1 : r]
        }
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
