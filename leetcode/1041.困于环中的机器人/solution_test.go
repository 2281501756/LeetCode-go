package leetcode

import (
	"testing"
)

func TestRobotBoundedInCircle(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isRobotBounded(instructions string) bool {
	x, y, dir := 0, 0, 0
	d := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := range instructions {
		if instructions[i] == 'G' {
			x += d[dir][0]
			y += d[dir][1]
		} else if instructions[i] == 'L' {
			dir += 3
			dir = dir % 4
		} else {
			dir++
			dir = dir % 4
		}

	}
	return x == 0 && y == 0 || dir != 0
}

//leetcode submit region end(Prohibit modification and deletion)
