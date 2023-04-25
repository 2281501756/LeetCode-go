package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var nums []int
	read := bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &n)
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(read, &nums[i])
	}
	res, l, r := 0, 0, 0
	mmap := map[int]bool{}
	for ; r < n; r++ {
		for mmap[nums[r]] == true && l < r {
			delete(mmap, nums[l])
			l++
		}
		mmap[nums[r]] = true
		res = max(res, r-l+1)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
