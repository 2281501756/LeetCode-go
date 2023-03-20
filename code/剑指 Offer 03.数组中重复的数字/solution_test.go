package leetcode

import (
	"testing"
)

func TestShuZuZhongZhongFuDeShuZiLcof(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func findRepeatNumber(nums []int) int {
	mmap := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
        if _, has := mmap[nums[i]]; has {
            return nums[i]
        }
        mmap[nums[i]] = true
	}
    return -1
}

//leetcode submit region end(Prohibit modification and deletion)
