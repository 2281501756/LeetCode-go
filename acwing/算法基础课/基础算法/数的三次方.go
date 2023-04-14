package main

import (
	"bufio"
	"fmt"
	"os"
)

var n float64

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &n)
	l, r := -100.0, 100.0
	for r-l > 1e-8 {
		mid := (l + r) / 2
		if mid*mid*mid > n {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Printf("%.6f", l)
}
