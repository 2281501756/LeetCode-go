package leetcode

import (
	"testing"
)

func TestLruCache(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	size       int
	capacity   int
	hash       map[int]*Node
	head, tail *Node
}
type Node struct {
	key       int
	value     int
	pre, next *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		hash:     map[int]*Node{},
		head:     &Node{0, 0, nil, nil},
		tail:     &Node{0, 0, nil, nil},
	}
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
		item.value = value
		this.moveToHead(item)
	} else {
		t := &Node{key: key, value: value}
		this.hash[key] = t
		this.addHead(t)
		this.size++
		if this.size > this.capacity {
			t := this.removeTail()
			delete(this.hash, t.key)
			this.size--
		}
	}

}
func (this *LRUCache) moveToHead(node *Node) {
	this.remove(node)
	this.addHead(node)
}
func (this *LRUCache) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (this *LRUCache) addHead(node *Node) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}
func (this *LRUCache) removeTail() *Node {
	t := this.tail.pre
	t.pre.next = t.next
	t.next.pre = t.pre
	return t
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
