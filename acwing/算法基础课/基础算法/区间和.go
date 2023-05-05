package main

import (
	"bufio"
	"fmt"
	"os"
)

func quick_sort3(arr []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	mid := arr[(i+j)>>1]
	for i < j {
		for {
			i++
			if arr[i] >= mid {
				break
			}

		}
		for {
			j--
			if arr[j] <= mid {
				break
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	quick_sort3(arr, l, j)
	quick_sort3(arr, j+1, r)
}
func find(arr []int, k int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := (l + r) >> 1
		if arr[mid] >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l + 1
}

func main() {
	const N = 300010
	var n, m int
	read := bufio.NewReader(os.Stdin)
	all, add, query := make([]int, 0), make([][]int, 0), make([][]int, 0)
	var s, a [N]int
	fmt.Fscan(read, &n, &m)

	var k, v, l, r int
	for i := 0; i < n; i++ {
		fmt.Fscan(read, &k, &v)
		all = append(all, k)
		add = append(add, []int{k, v})
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(read, &l, &r)
		all = append(all, []int{l, r}...)
		query = append(query, []int{l, r})
	}
	quick_sort3(all, 0, len(all)-1)
	i := 1
	for j := 1; j < len(all); j++ {
		if all[j] != all[i-1] {
			all[i] = all[j]
			i++
		}
	}
	all = all[:i]
	for i = 0; i < len(add); i++ {
		index := find(all, add[i][0])
		a[index] += add[i][1]
	}
	for i = 1; i <= len(all); i++ {
		s[i] = s[i-1] + a[i]
	}
	for i = 0; i < len(query); i++ {
		left := find(all, query[i][0])
		right := find(all, query[i][1])
		fmt.Println(s[right] - s[left-1])
	}
}
