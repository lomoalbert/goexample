package main

import (
    "fmt"
    "strings"
)

/*no5.go*/


func main() {
    for _, i := range []int{15, 15, 3, 2, 2, 2, 3, 15, 15} {
        fmt.Println(strings.Repeat("*", i))
    }
}