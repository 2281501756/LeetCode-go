package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestMostCommonWord(t *testing.T) {
	a := []int{3, 2, 12, 4, 1, 2, 3}
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)
}

// leetcode submit region begin(Prohibit modification and deletion)
type Pair struct {
	Key   string
	value int
}
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}
func (p PairList) Less(i, j int) bool {
	return p[i].value < p[j].value
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func mostCommonWord(paragraph string, banned []string) string {
	var arr []string
	for i := 0; i < len(paragraph); i++ {
		s, j := "", 0
		for j = i; j < len(paragraph) && isWord(paragraph[j]); j++ {
			s = s + string(paragraph[j])
		}
		i = j
		if s != "" {
			s = strings.ToLower(s)
			arr = append(arr, s)
		}
	}

	mmap := make(map[string]int, 0)
	for i := range arr {
		mmap[arr[i]]++
	}
	sortarr := sortmmap(mmap)

	for _, v := range sortarr {
		if !hasItem(banned, v.Key) {
			return v.Key
		}
	}
	return ""
}
func isWord(s uint8) bool {
	if s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' {
		return true
	}
	return false
}
func sortmmap(mmap map[string]int) PairList {
	p := make(PairList, len(mmap))
	i := 0
	for k, v := range mmap {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p
}
func hasItem(arr []string, item string) bool {
	for i := range arr {
		if arr[i] == item {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
