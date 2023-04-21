package main

import "fmt"

func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true
}

func main() {
	var n, m string
	a, b := make([]int, 0), make([]int, 0)
	fmt.Scan(&n)
	fmt.Scan(&m)
	for i := len(n) - 1; i >= 0; i-- {
		a = append(a, int(n[i]-'0'))
	}
	for i := len(m) - 1; i >= 0; i-- {
		b = append(b, int(m[i]-'0'))
	}
	if cmp(a, b) {
	} else {
		a, b = b, a
		fmt.Print("-")
	}
	c := make([]int, 0)
	for i, t := 0, 0; i < len(a); i++ {
		t = a[i] - t
		if len(b) > i {
			t = t - b[i]
		}
		c = append(c, (t+10)%10)
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}
	for c[len(c)-1] == 0 && len(c) > 1 {
		c = c[:len(c)-1]
	}

	for i := len(c) - 1; i >= 0; i-- {
		fmt.Printf("%d", c[i])
	}
}
