package main

import "fmt"

type link1 struct {
	e   []int
	l   []int
	r   []int
	idx int
}

func (l *link1) init() {
	l.e = make([]int, 100010)
	l.l = make([]int, 100010)
	l.r = make([]int, 100010)
	l.r[0] = 1
	l.l[1] = 0
	l.idx = 2
}
func (l *link1) add(k, v int) {
	l.e[l.idx] = v
	l.l[l.idx] = k
	l.r[l.idx] = l.r[k]
	l.l[l.r[k]] = l.idx
	l.r[k] = l.idx
	l.idx++
}
func (l *link1) remove(k int) {
	l.r[l.l[k]] = l.r[k]
	l.l[l.r[k]] = l.l[k]
}

func main() {
	var n int
	fmt.Scan(&n)
	l := link1{}
	l.init()
	for i := 0; i < n; i++ {
		var op string
		fmt.Scan(&op)
		switch op {
		case "L":
			var x int
			fmt.Scan(&x)
			l.add(0, x)
		case "R":
			var x int
			fmt.Scan(&x)
			l.add(l.l[1], x)
		case "D":
			var k int
			fmt.Scan(&k)
			l.remove(k + 1)
		case "IL":
			var k, x int
			fmt.Scan(&k, &x)
			l.add(l.l[k+1], x)
		case "IR":
			var k, x int
			fmt.Scan(&k, &x)
			l.add(k+1, x)

		}
	}
	for i := l.r[0]; i != 1; i = l.r[i] {
		fmt.Printf("%d ", l.e[i])
	}
}
