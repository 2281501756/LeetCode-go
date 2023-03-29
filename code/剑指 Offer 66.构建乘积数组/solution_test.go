package leetcode

import (
	"testing"
)

func TestGouJianChengJiShuZuLcof(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func constructArr(a []int) []int {
    if len(a) == 0 {
        return []int{}
    }
	arr1, arr2, res := make([]int, len(a)), make([]int, len(a)), make([]int, len(a))
	arr1[0] = 1
	for i := 1; i < len(a); i++ {
		arr1[i] = arr1[i-1] * a[i-1]
	}
	arr2[len(arr2)-1] = 1
	for i := len(a) - 2; i >= 0; i-- {
		arr2[i] = arr2[i+1] * a[i+1]
	}
	for i := 0; i < len(a); i++ {
		res[i] = arr1[i] * arr2[i]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
