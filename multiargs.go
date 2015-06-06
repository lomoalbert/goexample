package main

import "fmt"

func getarg(args ...int) int {
    for i, j := range args {
        fmt.Println(i, j)
    }
    return len(args)
}

func main() {
    n := getarg(1, 2, 3, 4, 5, 6, 7, 8)
    fmt.Println("一共有", n, "个参数")
    e := []int{9, 8, 7, 6, 5, 4, 3, 2}
    n=getarg(e...)
    fmt.Println("一共有", n, "个参数")
}
