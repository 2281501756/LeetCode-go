package leetcode

import (
	"strings"
	"testing"
)

func TestXltzEq(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	mys := ""
	for i := range s {
		if isValid(s[i]) {
			mys += string(s[i])
		}
	}
	mys = strings.ToLower(mys)
	for i := 0; i < len(mys)/2; i++ {
		if mys[i] != mys[len(mys)-1-i] {
			return false
		}
	}
	return true
}
func isValid(s byte) bool {
	return (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9')
}

//leetcode submit region end(Prohibit modification and deletion)
