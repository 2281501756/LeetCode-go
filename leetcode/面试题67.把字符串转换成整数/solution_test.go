package leetcode

import (
	"math"
	"strings"
	"testing"
)

func TestBaZiFuChuanZhuanHuanChengZhengShuLcof(t *testing.T) {
	println(strToInt("+"))
}

// leetcode submit region begin(Prohibit modification and deletion)
func strToInt(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0

	}
	if str[0] != '-' && str[0] != '+' && !(str[0] >= '0' && str[0] <= '9') {
		return 0
	}
	flag := true
	if str[0] == '+' {
		str = str[1:]
	} else if str[0] == '-' {
		str = str[1:]
		flag = false
	}
	if len(str) == 0 {
		return 0

	}
	var num float64
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			num *= 10
			num += float64(str[i] - '0')
		} else {
			break
		}
	}
	if num > (math.Pow(2, 31)-1) && flag {
		return int(math.Pow(2, 31) - 1)
	}
	if num > math.Pow(2, 31) && !flag {
		return int(-math.Pow(2, 31))
	}
	if flag {
		return int(num)
	} else {
		return int(-num)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
