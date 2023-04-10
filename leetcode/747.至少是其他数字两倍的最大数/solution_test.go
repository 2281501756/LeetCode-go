package leetcode

import (
	"testing"
)

func TestLargestNumberAtLeastTwiceOfOthers(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func dominantIndex(nums []int) int {
	max, res := -1, -1
	for i := range nums {
		if max < nums[i] {
			max = nums[i]
            res = i
		}
	}

	for i := range nums {
		if nums[i] != max && nums[i]*2 > max {
			return -1
		}
	}
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
