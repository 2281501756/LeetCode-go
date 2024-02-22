package leetcode

import (
	"testing"
)

func TestVvJkup(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	status := make(map[int]bool)
	res := make([][]int, 0)
	var dfs func(int, []int)
	dfs = func(n int, arr []int) {
		if n >= len(nums) {
			res = append(res, arr)
			return
		}
		for _, num := range nums {
			if !status[num] {
				status[num] = true
				dfs(n+1, append(arr, num))
				status[num] = false
			}
		}
	}
	dfs(0, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
