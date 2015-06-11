package main

import (
    "time"
    "fmt"
)

func main(){
    //time.After返回一个chan,指定时间后通过chan传回即时时间
    fmt.Println(<- time.After(2*time.Second))
    //time.Sleep 阻塞指定时间
    time.Sleep(2 * time.Second)
    //time.Tick 类似time.After,但可以持续使用
    c := time.Tick(1 * time.Second)
    for now := range c {
        fmt.Printf("%v\n", now)
    }
}
