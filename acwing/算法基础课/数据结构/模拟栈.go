package main

import "fmt"

func main() {
	var n int
	var s []int
	var h, t int
	s = make([]int, 100010)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var op string
		fmt.Scan(&op)
		switch op {
		case "push":
			var x int
			fmt.Scan(&x)
			t++
			s[t] = x
		case "pop":
			//fmt.Println(s[t])
			t--
		case "empty":
			if h == t {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case "query":
			fmt.Println(s[t])
		}
	}
}
