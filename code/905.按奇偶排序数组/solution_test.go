package leetcode

import(
    "testing"
)

func TestSortArrayByParity(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParity(nums []int) []int {
    arr1, arr2 := make([]int, 0), make([]int, 0)
    for i := 0; i < len(nums); i ++ {
        if nums[i]% 2 == 0 {
            arr1 = append(arr1, nums[i])
        } else {
            arr2 = append(arr2, nums[i])
        }
    }
    return append(arr1, arr2...)
}
//leetcode submit region end(Prohibit modification and deletion)
