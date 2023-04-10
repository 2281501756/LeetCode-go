package leetcode

import (
	"testing"
)

func TestPeakIndexInAMountainArray(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func peakIndexInMountainArray(arr []int) int {
	arr = append([]int{-1}, arr...)
	arr = append(arr, -1)
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) >> 1
		if arr[mid-1] < arr[mid] && arr[mid+1] < arr[mid] {
			return mid - 1
		} else if arr[mid-1] < arr[mid] && arr[mid+1] > arr[mid] {
			l = mid + 1
		} else if arr[mid-1] > arr[mid] && arr[mid+1] < arr[mid] {
			r = mid - 1
		}
	}
	return l - 1
}

//leetcode submit region end(Prohibit modification and deletion)
