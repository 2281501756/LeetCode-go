package leetcode

import(
    "testing"
)

func TestTwoSum(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i ++ {
        for j := i + 1; j < len(nums); j ++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}
//leetcode submit region end(Prohibit modification and deletion)
