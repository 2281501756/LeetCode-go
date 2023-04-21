package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n, m int
	var nums []int

	read := bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &n, &m)
	nums = make([]int, n+2)
	for i := 1; i <= n; i++ {
		var t int
		fmt.Fscan(read, &t)
		insert(i, i, t, nums)
	}
	for m > 0 {
		var l, r, t int
		fmt.Fscan(read, &l, &r, &t)
		insert(l, r, t, nums)
		m--
	}
	for i := 1; i <= n; i++ {
		nums[i] += nums[i-1]
		fmt.Printf("%d ", nums[i])
	}

}

func insert(l, r, d int, nums []int) {
	nums[l] += d
	nums[r+1] -= d
}
