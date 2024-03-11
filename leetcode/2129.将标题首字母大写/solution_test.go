package leetcode

import (
	"strings"
	"testing"
)

func TestCapitalizeTheTitle(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func capitalizeTitle(title string) string {
	res := ""
	for i, j := 0, 0; i < len(title); i++ {
		if title[i] == ' ' {
			if i-j <= 2 {
				res += strings.ToLower(title[j : i+1])
			} else {
				res += strings.ToUpper(title[j : j+1])
				res += strings.ToLower(title[j+1 : i+1])
			}
			j = i + 1
		}
		if i == len(title)-1 && j <= i {
			if i-j <= 1 {
				res += strings.ToLower(title[j : i+1])
			} else {
				res += strings.ToUpper(title[j : j+1])
				res += strings.ToLower(title[j+1 : i+1])
			}
			j = i + 1
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
