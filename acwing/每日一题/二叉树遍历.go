// https://www.acwing.com/problem/content/description/3387/

package main

import "fmt"

var str string
var k int

func dfs() {
	if str[k] == '#' {
		k++
		return
	}
	t := str[k]
	k++
	dfs()
	fmt.Printf("%s ", string(t))
	dfs()
}

func main() {
	fmt.Scan(&str)
	dfs()
	return
}
