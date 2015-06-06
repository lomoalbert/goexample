package main

import (
    "time"
    "fmt"
)

func main() {
    // 首先，我们实现并执行一个匿名的超时等待函数
    timeout := make(chan bool, 1)
    ch := make(chan int, 1)
    go func() {
        time.Sleep(1e9) // 等待1秒钟
        timeout <- true
    }()

    // 然后我们把timeout这个channel利用起来
    select {
    case <-ch:
    // 从ch中读取到数据,但没有线程写入ch
    case <-timeout:
        fmt.Println("time out")
    // 一直没有从ch中读取到数据，但从timeout中读取到了数据
    }
}