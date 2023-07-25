package main

import "fmt"

func a() {
    c := make(chan int, 5)
    c <- 1
    v := <-c
}

func main() {
    s1 := []int{1, 2}
    s2 := s1
    s2 = append(s2, 3)
    add(s1)
    add(s2)
    fmt.Println(s1, s2, cap(s1), cap(s2))
}

func add(s []int) {
    s = append(s, 0)
    for i := range s {
        s[i]++
    }
}
