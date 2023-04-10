package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	arr := make([]int, n)
	st, tt := make([]int, n), 0
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		for tt > 0 && arr[i] <= st[tt-1] {
			tt--
		}
		if tt == 0 {
			fmt.Printf("%d ", -1)
		} else {
			fmt.Printf("%d ", st[tt-1])
		}
		st[tt] = arr[i]
		tt++
	}

}
