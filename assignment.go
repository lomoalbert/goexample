package main

import "fmt"

//多返回值赋值时,只要有一个变量是新变量,就可以:=
func main() {
    a:=10
    a, b := func() (x, y int) {
        return 1, 2
    }()
    fmt.Println(a, b)
}

