package leetcode

import(
    "strings"
    "testing"
)

func TestWordPattern(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
    m1 := map[string]byte{}
    m2 := map[byte]string{}
    words := strings.Split(s, " ")
    if len(words) != len(pattern) {
        return false
    }
    for i, word := range words {
        ch := pattern[i]
        if m1[word] > 0 && m1[word] != ch || m2[ch] != "" && m2[ch] != word {
            return false
        }
        m1[word] = ch
        m2[ch] = word
    }
    return true
}
//leetcode submit region end(Prohibit modification and deletion)
