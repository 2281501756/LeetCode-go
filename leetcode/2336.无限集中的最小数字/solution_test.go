package leetcode

import (
	"github.com/emirpasic/gods/maps/treemap"
	"testing"
)

func TestSmallestNumberInInfiniteSet(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
type SmallestInfiniteSet struct {
	m *treemap.Map
}

func Constructor() SmallestInfiniteSet {
	s := SmallestInfiniteSet{m: treemap.NewWithIntComparator()}
	for i := 1; i <= 1000; i++ {
		s.m.Put(i, nil)
	}
	return s
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	res, _ := this.m.Min()
	this.m.Remove(res.(int))
	return res.(int)
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	this.m.Put(num, nil)
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
//leetcode submit region end(Prohibit modification and deletion)
