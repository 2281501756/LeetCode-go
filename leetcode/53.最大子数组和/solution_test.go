package leetcode

import (
    "testing"
)

func TestMaximumSubarray(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
    dp := make([]int, len(nums))
    dp[0] = nums[0]
    res := dp[0]
    for i := 1; i < len(nums); i ++ {
        if nums[i] > dp[i - 1] + nums[i] {
            dp[i] = nums[i]
        } else {
            dp[i] = dp[i - 1] + nums[i]
        }
        if res < dp[i] {
            res = dp[i]
        }
    }
    return res
}

//leetcode submit region end(Prohibit modification and deletion)
