package leetcode

import (
	"math"
	"testing"
)

func TestLargestTriangleArea(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func largestTriangleArea(points [][]int) float64 {
	res := float64(-1)
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for z := j + 1; z < len(points); z++ {
				x1, y1 := points[i][0], points[i][1]
				x2, y2 := points[j][0], points[j][1]
				x3, y3 := points[z][0], points[z][1]
				res = max(res, math.Abs(float64(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2))/2))
			}
		}
	}
	return res
}
func max(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
