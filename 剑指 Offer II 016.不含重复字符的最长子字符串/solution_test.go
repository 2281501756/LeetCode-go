package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	hash := map[rune]int{}
	res := 0
	l := 0
	for _, v := range s {
		if _, ok := hash[v]; ok {
			for ok {
				delete(hash, rune(s[l]))
				l++
				_, ok = hash[v]
			}
		}
		hash[v] = 1
		res = max(res, len(hash))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWtcaE1(t *testing.T) {

}
