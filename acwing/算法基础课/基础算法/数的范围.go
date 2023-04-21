package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var n, m int
    var nums []int

    read := bufio.NewReader(os.Stdin)
    fmt.Fscanf(read, "%d %d", &n, &m)
    nums = make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(read, &nums[i])
    }
    for m > 0 {
        var t int
        fmt.Fscan(read, &t)
        l, r := 0, n-1
        for l < r {
            mid := (l + r) >> 1
            if nums[mid] >= t {
                r = mid
            } else {
                l = mid + 1
            }
        }
        if nums[l] == t {
            fmt.Printf("%d ", l)
        } else {
            fmt.Printf("-1 ")
        }
        l, r = 0, n-1
        for l < r {
            mid := (l + r + 1) >> 1
            if nums[mid] > t {
                r = mid - 1
            } else {
                l = mid
            }
        }
        if nums[l] == t {
            fmt.Printf("%d\n", l)
        } else {
            fmt.Printf("-1\n")
        }
        m--
    }
}
