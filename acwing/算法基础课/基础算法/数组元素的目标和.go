package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, k int
	var arr1, arr2 []int
	read := bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &n, &m, &k)
	arr1, arr2 = make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(read, &arr1[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(read, &arr2[i])
	}
	for i, j := 0, m-1; i < n; i++ {
		for ; j >= 0 && arr1[i]+arr2[j] >= k; j-- {
			if arr1[i]+arr2[j] == k {
				fmt.Printf("%d %d", i, j)
			}
		}
	}
}
