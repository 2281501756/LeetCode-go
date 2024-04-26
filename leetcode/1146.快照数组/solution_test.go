package leetcode

import (
	"sort"
	"testing"
)

func TestSnapshotArray(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
type SnapshotArray struct {
	snapCnt int
	data    [][][2]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		snapCnt: 0,
		data:    make([][][2]int, length),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.data[index] = append(this.data[index], [2]int{this.snapCnt, val})
}

func (this *SnapshotArray) Snap() int {
	this.snapCnt++
	return this.snapCnt - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	x := sort.Search(len(this.data[index]), func(i int) bool {
		return this.data[index][i][0] > snap_id
	})
	if x == 0 {
		return 0
	}

	return this.data[index][x-1][1]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
//leetcode submit region end(Prohibit modification and deletion)
