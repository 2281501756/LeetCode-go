package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(read, &n)
	for t := 1; t <= n; t++ {
		var s1, s2 string
		fmt.Fscan(read, &s1, &s2)
		i, j := 0, 0
		for ; j < len(s2); j++ {
			if s1[i] == s2[j] {
				i++
			}
			if i >= len(s1) {
				break
			}
		}
		if i == len(s1) {
			fmt.Printf("Case #%d: %d\n", t, len(s2)-len(s1))
		} else {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", t)
		}

	}
}
