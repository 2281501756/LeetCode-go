package leetcode

import (
	"testing"
)

func TestQ91FMA(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func lenLongestFibSubseq(arr []int) int {
	hash := map[int]int{}
	for i, item := range arr {
		hash[item] = i
	}
	n := len(arr)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			x := arr[i] - arr[j]
			f[i][j] = 2
			if k, ok := hash[x]; x < arr[j] && ok {
				f[i][j] = max(f[i][j], f[j][k]+1)
				res = max(res, f[i][j])
			}
		}
	}
	if res < 3 {
		res = 0
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
