package leetcode

import (
	"sort"
	"testing"
)

func TestSmallestNumberInInfiniteSet(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
type SmallestInfiniteSet struct {
	increaseId int
	arr        []int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		increaseId: 1,
		arr:        []int{},
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.arr) == 0 {
		this.increaseId++
		return this.increaseId - 1
	} else {
		res := this.arr[0]
		this.arr = this.arr[1:]
		return res
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.increaseId {
		return
	}
	this.arr = append(this.arr, num)
	unique_sort(this.arr)
}

func unique_sort(slice []int) []int {
	res := make([]int, 0)
	m := make(map[int]bool, len(slice))
	for _, e := range slice {
		if m[e] == false {
			m[e] = true
			res = append(res, e)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
//leetcode submit region end(Prohibit modification and deletion)
