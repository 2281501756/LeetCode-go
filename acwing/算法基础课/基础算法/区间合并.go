package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	var res, sortArray [][2]int
	read := bufio.NewReader(os.Stdin)
	fmt.Fscan(read, &n)
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(read, &l, &r)
		sortArray = append(sortArray, [2]int{l, r})
	}
	sort.Slice(sortArray, func(i, j int) bool {
		return sortArray[i][0] < sortArray[j][0] || (sortArray[i][0] == sortArray[j][0] && sortArray[i][1] < sortArray[j][1])
	})
	res = append(res, sortArray[0])
	for i := 1; i < len(sortArray); i++ {
		if sortArray[i][0] > res[len(res)-1][1] {
			res = append(res, sortArray[i])
		} else {
			res[len(res)-1][1] = max1(res[len(res)-1][1], sortArray[i][1])
		}
	}
	fmt.Println(len(res))
}
func max1(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
