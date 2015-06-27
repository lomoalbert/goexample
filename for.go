package main

import "fmt"

func main() {
    // 初始化i为0；每一次循环之后检查i是否小于5；让i加1
    for i := 0; i < 5; i++ {
        fmt.Println("i现在的值为：", i)
    }

    // for ;; 定义的变量仅限循环中使用
    for i := 0; i < 5; {
        fmt.Println("i现在的值为：", i)
        i+=1
    }
    //fmt.Println(i) //undefined: i

    // for 以外定义的变量正常使用
    i := 0
    for ; i<5; i++ {
        fmt.Println("i现在的值为：", i)
    }
    fmt.Println("i最终的值为：", i)

    for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s + "a" {
        fmt.Printf("i=%d, j=%d, s=%s\n", i, j, s)
    }
}
