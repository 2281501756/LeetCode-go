package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	n, t := 0, 0
	finalDay, res, store := 0, 0, 0
	fmt.Fscan(read, &n, &t)
	for n > 0 {
		n--
		di, ti := 0, 0
		fmt.Fscan(read, &di, &ti)
		if di-finalDay >= store {
			res += store
			store = ti
		} else {
			res += di - finalDay
			store = store - (di - finalDay) + ti
		}
		finalDay = di
	}
	if store > 0 {
		if finalDay+store > t {
			res += t - finalDay + 1
		} else {
			res += store
		}
	}
	fmt.Println(res)
}
