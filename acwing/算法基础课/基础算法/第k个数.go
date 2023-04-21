package main

import (
	"bufio"
	"fmt"
	"os"
)

var k int

func main() {

	var n int
	var nums []int
	read := bufio.NewReader(os.Stdin)
	fmt.Fscanf(read, "%d %d", &n, &k)
	nums = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(read, &nums[i])
	}
	quick_sort1(nums, 0, n-1)
	fmt.Printf("%d", nums[k-1])
}

func quick_sort1(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	mid := nums[(i+j)>>1]
	for i < j {
		for {
			i++
			if nums[i] >= mid {
				break
			}
		}
		for {
			j--
			if nums[j] <= mid {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if j >= k-1 {
		quick_sort(nums, r, j)
	} else {
		quick_sort(nums, j+1, r)
	}
}
