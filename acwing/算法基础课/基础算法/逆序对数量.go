package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 10000

var temp, q [N]int

func merge_sort1(l, r int) int {
	if l >= r {
		return 0
	}
	mid := (l + r) >> 1
	res := merge_sort1(l, mid) + merge_sort1(mid+1, r)
	i, j, k := l, mid+1, 0
	for i <= mid && j <= r {
		if q[i] < q[j] {
			temp[k] = q[i]
			k++
			i++
		} else {
			temp[k] = q[j]
			res += mid - i + 1
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = q[j]
		k++
		j++
	}
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = temp[j]
	}
	return res
}

func main() {
	var n int
	read := bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(read, &q[i])
	}
	fmt.Print(merge_sort1(0, n-1))
}
