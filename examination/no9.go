package main

import (
    "fmt"
)

/*程序 09：象棋棋盘

题目：要求输出国际象棋棋盘。
1.程序分析：用 i 控制行， j 来控制列，根据i+j的和的变化来控制输出黑方格，还是白方格。
2.程序源代码：*/


func main() {
    for i := 0; i<8; i++ {
        for j := 0; j<8; j++ {
            if (i+j)%2==0 {
                fmt.Printf("■")
            }else {
                fmt.Printf("□")
            }
        }
        fmt.Println("")
    }
}