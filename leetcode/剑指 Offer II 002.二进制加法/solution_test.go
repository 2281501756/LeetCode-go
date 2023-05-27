package leetcode

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestJFETK5(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	res := make([]string, 0)
	t := 0
	for z, i, j := 0, len(a)-1, len(b)-1; z < max(len(a), len(b)); z++ {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			n2 += int(b[j] - '0')
			j--
		}
		res = append(res, strconv.Itoa((n1+n2+t)%2))
		t = (n1 + n2 + t) / 2
	}
	if t == 1 {
		res = append(res, strconv.Itoa(1))
	}
	sort.SliceStable(res, func(i, j int) bool {
		return true
	})
	return strings.Join(res, "")
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
