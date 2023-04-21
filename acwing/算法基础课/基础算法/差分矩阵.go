package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n, m, k int
	line := readLine(scanner)
	n, m, k = line[0], line[1], line[2]
	a, b := make([][]int, n+2), make([][]int, n+2)
	for i := range a {
		a[i] = make([]int, m+2)
		b[i] = make([]int, m+2)
	}
	for i := 1; i <= n; i++ {
		line := readLine(scanner)
		for j := 1; j <= m; j++ {
			a[i][j] = line[j-1]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			insert1(b, i, j, i, j, a[i][j])
		}
	}

	for k > 0 {
		line = readLine(scanner)
		insert1(b, line[0], line[1], line[2], line[3], line[4])
		k--
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			b[i][j] += b[i-1][j] + b[i][j-1] - b[i-1][j-1]
		}
	}
	for i := 1; i <= n; i++ {
		row := make([]string, m)
		for j := 1; j <= m; j++ {
			row[j-1] = strconv.Itoa(b[i][j])
		}
		fmt.Println(strings.Join(row, " "))
	}
}

func insert1(arr [][]int, x1, y1, x2, y2, k int) {
	arr[x1][y1] += k
	arr[x1][y2+1] -= k
	arr[x2+1][y1] -= k
	arr[x2+1][y2+1] += k
}

func readLine(scanner *bufio.Scanner) []int {
	scanner.Scan()
	p := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	line := make([]int, len(p))
	for i, item := range p {
		line[i], _ = strconv.Atoi(item)
	}
	return line
}
