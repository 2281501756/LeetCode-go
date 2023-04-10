package leetcode

import(
    "testing"
)

func TestRemoveElement(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
    i, j := 0, 0
    for i, j = 0, 0; i < len(nums); i ++ {
        if nums[i] != val {
            nums[j] = nums[i]
            j++
        }
    }
    return j
}
//leetcode submit region end(Prohibit modification and deletion)
