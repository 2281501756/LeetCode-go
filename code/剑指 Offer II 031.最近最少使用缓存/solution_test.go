package leetcode

import (
	"testing"
)

func TestOrIXps(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	size       int
	capacity   int
	hash       map[int]*Node
	head, tail *Node
}
type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{hash: map[int]*Node{}, size: 0, capacity: capacity, head: &Node{0, 0, nil, nil}, tail: &Node{0, 0, nil, nil}}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	item, has := this.hash[key]
	if !has {
		return -1
	}
	this.moveToHead(item)
	return item.value
}

func (this *LRUCache) Put(key int, value int) {
	item, has := this.hash[key]
	if has {
		this.moveToHead(item)
		item.value = value
	} else {
		item := &Node{key, value, nil, nil}
		this.addHead(item)
		this.hash[key] = item
		this.size++
		if this.size > this.capacity {
			t := this.removeTail()
			delete(this.hash, t.key)
			this.size--
		}
	}
}
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addHead(node)
}
func (this *LRUCache) removeNode(item *Node) {
	item.pre.next = item.next
	item.next.pre = item.pre
}
func (this *LRUCache) addHead(item *Node) {
	item.pre = this.head
	item.next = this.head.next
	this.head.next.pre = item
	this.head.next = item
}
func (this *LRUCache) removeTail() *Node {
	t := this.tail.pre
	this.tail.pre = this.tail.pre.pre
	this.tail.pre.next = this.tail
	return t
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
