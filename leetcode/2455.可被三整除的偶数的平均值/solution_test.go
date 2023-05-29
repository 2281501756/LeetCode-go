package leetcode

import (
    "testing"
)

func TestAverageValueOfEvenNumbersThatAreDivisibleByThree(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func averageValue(nums []int) int {
    res, n := 0, 0
    for i := 0; i < len(nums); i ++ {
        if nums[i] % 3 == 0 && nums[i] % 2 == 0{
            res += nums[i]
            n++
        }
    }
    if n == 0 {
        return 0
    }
    return  res / n
}
//leetcode submit region end(Prohibit modification and deletion)
