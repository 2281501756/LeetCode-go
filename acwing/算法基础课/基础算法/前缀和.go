package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var nums []int

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &n, &m)
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(read, &nums[i])
	}
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]
	}
	for m > 0 {
		var l, r int
		fmt.Fscan(read, &l, &r)
		fmt.Println(s[r] - s[l-1])
		m--
	}
}
