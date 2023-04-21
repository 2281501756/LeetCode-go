package main

import "fmt"

func main() {
	var (
		a []int
		b int
		c []int
		t int
		s string
	)
	fmt.Scan(&s)
	fmt.Scan(&b)
	for i := len(s) - 1; i >= 0; i-- {
		a = append(a, int(s[i]-'0'))
	}
	for i := 0; i < len(a) || t != 0; i++ {
		if i < len(a) {
			t += a[i] * b
		}
		c = append(c, t%10)
		t = t / 10
	}
	for c[len(c)-1] == 0 && len(c) > 1 {
		c = c[:len(c)-1]
	}
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Printf("%d", c[i])
	}
}
