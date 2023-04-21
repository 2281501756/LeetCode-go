package main

import "fmt"

func main() {
	var (
		a []int
		b int
		c []int
		s string
		t int
	)
	fmt.Scan(&s)
	fmt.Scan(&b)
	for i := 0; i < len(s); i++ {
		a = append(a, int(s[i]-'0'))
	}
	for i := 0; i < len(a); i++ {
		t = t*10 + a[i]
		c = append(c, t/b)
		t = t % b
	}
	for len(c) > 1 && c[0] == 0 {
		c = c[1:len(c)]
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("%d", c[i])
	}
	fmt.Printf("\n%d", t)
}
