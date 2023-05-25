package leetcode

import (
	"testing"
)

func TestOddStringDifference(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func oddString(words []string) string {
    w1, w2 := getArr(words[0]), getArr(words[1])
    if arrIsSimilar(w1, w2) {
        for i := 2; i < len(words); i ++ {
            if !arrIsSimilar(w1, getArr(words[i])) {
                return words[i]
            }
        }
        return words[0]
    } else {
        w3 := getArr(words[2])
        if arrIsSimilar(w1, w3) {
            return words[1]
        } else {
            return words[0]
        }
    }
}
func getArr(word string) []int {
	res := make([]int, 0)
	for i := 0; i < len(word)-1; i++ {
		res = append(res, int(word[i+1]-word[i]))
	}
	return res
}
func arrIsSimilar(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
