package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var n int
    read := bufio.NewReader(os.Stdin)
    fmt.Fscan(read, &n)
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(read, &nums[i])
        fmt.Printf("%d ", lowbiy(nums[i]))
    }

}

func lowbiy(n int) (res int) {
    for n > 0 {
        n -= n & -n
        res++
    }
    return res
}
