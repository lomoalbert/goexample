package main

import (
    "fmt"
)

/*程序 08：9*9 口诀

题目：输出 9*9 口诀。
1.程序分析：分行与列考虑，共 9 行 9 列，i 控制行，j 控制列。
2.程序源代码：*/


func main() {
    for i:=1;i<10;i++{
        for j:=1;j<=i;j++{
            fmt.Printf("%d*%d=%d\t",j,i,i*j)
        }
        fmt.Println("")
    }
}