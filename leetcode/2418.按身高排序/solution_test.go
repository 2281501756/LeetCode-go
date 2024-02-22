package leetcode

import (
	"testing"
)

func TestSortThePeople(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func sortPeople(names []string, heights []int) []string {
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights)-i-1; j++ {
			if heights[j] < heights[j+1] {
				heights[j], heights[j+1] = heights[j+1], heights[j]
				names[j], names[j+1] = names[j+1], names[j]
			}
		}
	}
	return names
}

//leetcode submit region end(Prohibit modification and deletion)
