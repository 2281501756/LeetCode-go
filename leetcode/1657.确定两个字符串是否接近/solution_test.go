package leetcode

import (
	"sort"
	"testing"
)

func TestDetermineIfTwoStringsAreClose(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func closeStrings(word1 string, word2 string) bool {
	m1, m2 := map[byte]int{}, map[byte]int{}
	if len(word1) != len(word2) {
		return false
	}
	for i := 0; i < len(word1); i++ {
		m1[word1[i]]++
	}
	for i := 0; i < len(word2); i++ {
		m2[word2[i]]++
	}
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
	}
	arr1, arr2 := make([]int, len(m1)), make([]int, len(m1))
	for _, v := range m1 {
		arr1 = append(arr1, v)
	}
	for _, v := range m2 {
		arr2 = append(arr2, v)
	}
	sort.Ints(arr1[:])
	sort.Ints(arr2[:])
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
