package leetcode

import (
	"testing"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 获取总长度
	total := len(nums1) + len(nums2)
	if total%2 == 0 {
		// 获取中间的两个值
		left, right := find(nums1, 0, nums2, 0, total/2), find(nums1, 0, nums2, 0, total/2+1)
		return (float64(left) + float64(right)) / 2
	} else {
		// 获取中间值
		return float64(find(nums1, 0, nums2, 0, total/2+1))
	}
}

func find(nums1 []int, i int, nums2 []int, j int, k int) int {
	// 如果num1剩余的大小更多我们就交换下
	if len(nums1)-i > len(nums2)-j {
		return find(nums2, j, nums1, i, k)
	}
	// 如果num1用完了直接输出num2中对应的数即可
	if i == len(nums1) {
		return nums2[j+k-1]
	}
	// 如果k==1表示在num1和num2中输出最小的即可
	if k == 1 {
		return min(nums1[i], nums2[j])
	}
	// 获取中间值并比较，将小的一方删除掉部分
	si, sj := min(len(nums1), i+k/2), j+k-k/2
	if nums1[si-1] > nums2[sj-1] {
		return find(nums1, i, nums2, sj, k-(sj-j))
	} else {
		return find(nums1, si, nums2, j, k-(si-i))
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
