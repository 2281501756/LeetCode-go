package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	read := bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &n, &m)
	arr1, arr2 := make([]int, n), make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(read, &arr1[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(read, &arr2[i])
	}
	for i, j := 0, 0; i < m; i++ {
		if arr2[i] == arr1[j] {
			j++
			if j == n {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
