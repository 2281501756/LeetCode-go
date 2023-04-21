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
	merge_sort(nums, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", nums[i])
	}
}

func merge_sort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	merge_sort(nums, l, mid)
	merge_sort(nums, mid+1, r)
	arr := make([]int, 0)
	i, j := l, mid+1
	for i <= mid && j <= r {
		if nums[i] < nums[j] {
			arr = append(arr, nums[i])
			i++
		} else {
			arr = append(arr, nums[j])
			j++
		}
	}
	for i <= mid {
		arr = append(arr, nums[i])
		i++
	}
	for j <= r {
		arr = append(arr, nums[j])
		j++
	}
	for i := 0; i < len(arr); i++ {
		nums[l+i] = arr[i]
	}
}
