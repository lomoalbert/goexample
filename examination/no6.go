package main

import (
    "fmt"
    "strings"
)

/*程序 06：打印图形

题目：用*号输出字母 C 的图案。
1.程序分析：可先用'*'号在纸上写出字母 C，再分行输出。
2.程序源代码：*/


func main() {
    for _, i := range []int{15, 15, 3, 2, 2, 2, 3, 15, 15} {
        fmt.Println(strings.Repeat("*", i))
    }
}