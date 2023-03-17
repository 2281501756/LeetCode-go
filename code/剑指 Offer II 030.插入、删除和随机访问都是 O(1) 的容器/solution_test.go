package leetcode

import (
	"math/rand"
	"testing"
)

func TestFortPu(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
type RandomizedSet struct {
	indexList []int
	hash      map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, has := this.hash[val]
	if has {
		return false
	}
	this.hash[val] = len(this.indexList)
	this.indexList = append(this.indexList, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	id, has := this.hash[val]
	if !has {
		return false
	}
	last := len(this.indexList) - 1
	this.indexList[id] = this.indexList[last]
	this.hash[this.indexList[id]] = id
	this.indexList = this.indexList[:last]
	delete(this.hash, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.indexList[rand.Intn(len(this.indexList))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)
