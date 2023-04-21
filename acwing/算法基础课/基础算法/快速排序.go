package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var n int
    var nums []int

    read := bufio.NewReader(os.Stdin)
    fmt.Fscan(read, &n)
    nums = make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(read, &nums[i])
    }
    quick_sort(nums, 0, n-1)
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", nums[i])
    }
}

func quick_sort(nums []int, l int, r int) {
    if l >= r {
        return
    }
    i, j := l-1, r+1
    mid := nums[(i+j)>>1]
    for i < j {
        for {
            i++
            if nums[i] >= mid {
                break
            }
        }
        for {
            j--
            if nums[j] <= mid {
                break
            }
        }
        if i < j {
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    quick_sort(nums, l, j)
    quick_sort(nums, j+1, r)
}
