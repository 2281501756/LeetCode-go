package leetcode

import (
	"testing"
)

func TestAmbiguousCoordinates(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func ambiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]
	res := make([]string, 0)
	for t := 1; t < len(s); t++ {
		lt := getPos(s[:t])
		if len(lt) == 0 {
			continue
		}
		rt := getPos(s[t:])
		if len(rt) == 0 {
			continue
		}
		for _, i := range lt {
			for _, j := range rt {
				res = append(res, "("+i+", "+j+")")
			}
		}
	}
	return res
}

func getPos(s string) (pos []string) {
	if s[0] != '0' || s == "0" {
		pos = append(pos, s)
	}
	for i := 1; i < len(s); i++ {
		if i != 1 && s[0] == '0' || s[len(s)-1] == '0' {
			continue
		}
		pos = append(pos, s[:i]+"."+s[i:])
	}
	return pos
}

//leetcode submit region end(Prohibit modification and deletion)
