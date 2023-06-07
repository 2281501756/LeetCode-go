package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1010

var n, m int
var v, w []int
var f []int

func init() {
	v = make([]int, N)
	w = make([]int, N)
	f = make([]int, N)
}

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(read, &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i-1]; j-- {
			f[j] = max(f[j], f[j-v[i-1]]+w[i-1])
		}
	}
	fmt.Println(f[m])
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
