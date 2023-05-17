package main

import "fmt"

const N = 100010

type link struct {
	h     int
	e     []int
	ne    []int
	index int
}

func (l *link) init() {
	l.h = -1
	l.e = make([]int, N)
	l.ne = make([]int, N)
	l.index++
}
func (l *link) addToHead(v int) {
	l.e[l.index] = v
	l.ne[l.index] = l.h
	l.h = l.index
	l.index++
}
func (l *link) add(k, v int) {
	l.e[l.index] = v
	l.ne[l.index] = l.ne[k]
	l.ne[k] = l.index
	l.index++
}
func (l *link) remove(k int) {
	l.ne[k] = l.ne[l.ne[k]]
}
func (l *link) traverse() []int {
	res := make([]int, 0)
	for i := l.h; i != -1; i = l.ne[i] {
		res = append(res, l.e[i])
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	l := link{}
	l.init()
	for i := 0; i < n; i++ {
		var op string
		fmt.Scan(&op)
		switch op {
		case "H":
			var k int
			fmt.Scan(&k)
			l.addToHead(k)
		case "D":
			var k int
			fmt.Scan(&k)
			if k == 0 {
				l.h = l.ne[l.h]
			} else {
				l.remove(k)
			}
		case "I":
			var k, v int
			fmt.Scan(&k, &v)
			l.add(k, v)
		}
	}
	res := l.traverse()
	for _, v := range res {
		fmt.Printf("%d ", v)
	}
}
