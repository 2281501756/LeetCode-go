package leetcode

import(
    "strings"
    "testing"
)

func TestShortestCompletingWord(t *testing.T){
    
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestCompletingWord(licensePlate string, words []string) string {
    licensePlate = strings.ToLower(licensePlate)
    map1 := make(map[uint8]int, 0)
    for i := range licensePlate {
        s := licensePlate[i]
        if s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z'{
            map1[licensePlate[i]] ++
        }
    }
    res, maxleng := "", -1
    for i := range words {
        map2 := make(map[uint8]int, 0)
        for j := range words[i] {
            map2[words[i][j]] ++
        }
        flag := true
        for k, _ := range map1 {
            if map1[k] > map2[k] {
                flag = false
                break
            }
        }
        if !flag {
            continue
        }
        if maxleng == -1 || len(words[i]) < maxleng {
            res = words[i]
            maxleng = len(words[i])
        }
    }
    return res
}
//leetcode submit region end(Prohibit modification and deletion)
