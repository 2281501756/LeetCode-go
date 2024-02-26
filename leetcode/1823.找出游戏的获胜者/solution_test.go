package leetcode

import (
	"fmt"
	"testing"
)

func TestFindTheWinnerOfTheCircularGame(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findTheWinner(n int, k int) int {
    if k == 1 {
        return n
    }
	var dfs func(int, []int) int

	dfs = func(start int, ban []int) int {
		fmt.Println(start, ban)
		if len(ban) == n-1 {
			for contain(ban, start%n) {
				start++
			}
			return start % n
		}
		for i := 0; i < k; {
			for contain(ban, start%n) {
				start++
			}
			i++
			start++
			if i == k-1 {
				for contain(ban, start%n) {
					start++
				}
				ban = append(ban, start%n)
				return dfs((start+1)%n, ban)
			}
		}
		return 0
	}
	return dfs(0, []int{}) + 1
}

func contain(arr []int, n int) bool {
	for _, a := range arr {
		if a == n {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
