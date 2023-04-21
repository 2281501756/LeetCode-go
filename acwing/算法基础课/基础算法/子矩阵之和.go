package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	var n, m, k int
	fmt.Fscan(read, &n, &m, &k)
	s := make([][]int, n+1)
	for i := range s {
		s[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(read, &s[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] += s[i-1][j] + s[i][j-1] - s[i-1][j-1]
		}
	}
	for k > 0 {
		x1, y1, x2, y2 := 0, 0, 0, 0
		fmt.Fscan(read, &x1, &y1, &x2, &y2)
		fmt.Println(s[x2][y2] - s[x1-1][y2] - s[x2][y1-1] + s[x1-1][y1-1])
		k--
	}
}
