package main

import "fmt"

//普通用法 for init; condition;post {}
//while用法 for condition {}
//死循环 for {}
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
    //包含的while的功能
    for i<5 {
        i+=1
        fmt.Println("i现在的值为：", i)
    }
    fmt.Println("i最终的值为：", i)

    //for 循环
    for {
        if i>7 {
            fmt.Println("i最终的值为：", i)
            break //跳出循环
        }else {
            fmt.Println("i现在的值为：", i)
            i+=1
            continue //结束本次循环,直接转到下次循环
            i+=10 //continue 后的代码不会执行
        }
    }

    // 多变量条件
    for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s + "a" {
        fmt.Printf("i=%d, j=%d, s=%s\n", i, j, s)
    }
}
