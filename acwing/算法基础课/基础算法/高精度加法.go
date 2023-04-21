package main

import "fmt"

func main() {
	var n, m string
	a, b := make([]int, 0), make([]int, 0)
	fmt.Scanf("%s %s", &n, &m)
	for i := len(n) - 1; i >= 0; i-- {
		a = append(a, int(n[i]-'0'))
	}
	for i := len(m) - 1; i >= 0; i-- {
		b = append(b, int(m[i]-'0'))
	}
	c := make([]int, 0)
	t := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) {
			t += a[i]
		}
		if i < len(b) {
			t += b[i]
		}
		c = append(c, t%10)
		t = t / 10
	}
	if t != 0 {
		c = append(c, t)
	}
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Printf("%d", c[i])
	}
}
