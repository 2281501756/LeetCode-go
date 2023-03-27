package leetcode

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	res := isAnagram("abc", "bcd")
	fmt.Println(res)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mmap := make(map[int32]int)
	for _, i := range s {
		mmap[i]++
	}
	for _, i := range t {
		if n, has := mmap[i]; has {
			mmap[i] = n - 1
			if n-1 == 0 {
				delete(mmap, i)
			}
		} else {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
