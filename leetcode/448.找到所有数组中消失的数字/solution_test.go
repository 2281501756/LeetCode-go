package leetcode

import (
	"testing"
)

func TestFindAllNumbersDisappearedInAnArray(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	hash := make(map[int]bool, len(nums))
	for i := range nums {
		hash[nums[i]] = true
	}
	res := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		if _, ok := hash[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
