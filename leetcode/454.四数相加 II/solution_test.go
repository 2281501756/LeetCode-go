package leetcode

import (
    "testing"
)

func TestFourSumIi(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    mmap := map[int]int{}
    res := 0
    for _, a := range nums1 {
        for _, b := range nums2 {
            mmap[a + b] ++
        }
    }
    for _, c := range nums3 {
        for _, d := range nums4 {
            res += mmap[-c-d]
        }
    }
    return  res
}

//leetcode submit region end(Prohibit modification and deletion)
