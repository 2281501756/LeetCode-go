package main

import "fmt"

const N = 110

var (
	g, d [N][N]int
	n, m int
)

type point = struct {
	x, y int
}

func main() {
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&g[i][j])
			d[i][j] = -1
		}
	}
	fmt.Println(bfs())
}
func bfs() int {
	queue := make([]point, 0)
	queue = append(queue, point{x: 0, y: 0})
	d[0][0] = 0
	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x, y := top.x+dx[i], top.y+dy[i]
			if x >= 0 && x < n && y >= 0 && y < m && d[x][y] == -1 && g[x][y] == 0 {
				d[x][y] = d[top.x][top.y] + 1
				queue = append(queue, point{x: x, y: y})
			}
		}
	}
	return d[n-1][m-1]
}
