package leetcode

import (
	"fmt"
	"testing"
)

func TestValidSudoku(t *testing.T) {

}

// leetcode submit region begin(Prohibit modification and deletion)
func isValidSudoku(board [][]byte) bool {
	mmap := map[byte]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if mmap[board[i][j]] {
				return false
			}
			mmap[board[i][j]] = true
		}
		clearMap(mmap)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if mmap[board[j][i]] {
				return false
			}
			mmap[board[j][i]] = true
		}
		clearMap(mmap)
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
                    if board[i+x][j+y] == '.' {
						continue
					}
					if mmap[board[i+x][j+y]] {
						return false
					}
					mmap[board[i+x][j+y]] = true
				}
			}
			fmt.Print("\n")
			clearMap(mmap)
		}
	}
	return true
}

func clearMap(m map[byte]bool) {
	for k := range m {
		m[k] = false
	}
}

//leetcode submit region end(Prohibit modification and deletion)
