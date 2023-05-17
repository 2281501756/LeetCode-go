package leetcode

import (
    "testing"
)

func TestTheEmployeeThatWorkedOnTheLongestTask(t *testing.T) {

}

//leetcode submit region begin(Prohibit modification and deletion)
func hardestWorker(n int, logs [][]int) int {
    index := 0
    for i := 1; i < len(logs); i ++ {
        if index == 0{
            if logs[i][1] - logs[i - 1][1] > logs[0][1] {
                index = i
            } else if logs[i][1] - logs[i - 1][1] == logs[0][1] && logs[i][0] < logs[0][0] {
                index = i
            }
        } else {
            if logs[i][1] - logs[i - 1][1] > logs[index][1] - logs[index - 1][1] {
                index = i
            } else if logs[i][1] - logs[i - 1][1] == logs[index][1] - logs[index - 1][1] && logs[i][0] < logs[index][0] {
                index = i
            }
        }
    }
    if index == 0 {
        return logs[0][0]
    } else {
        return logs[index][0]
    }
}

//leetcode submit region end(Prohibit modification and deletion)
