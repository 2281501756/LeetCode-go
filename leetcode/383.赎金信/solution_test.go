package leetcode

import (
	"testing"
)

func TestRansomNote(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	cnt := map[int32]int{}
	for _, ch := range magazine {
		cnt[ch]++
	}
	for _, ch := range ransomNote {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
