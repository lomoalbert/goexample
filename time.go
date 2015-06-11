package main

import (
    "time"
    "fmt"
)
/*
type Duration int64
以int64表示两个时间点之间的间隔（纳秒数），最多表示290年

常量

Nanosecond Duration = 1
Microsecond = 1000 * Nanosecond     (1e3)
Millissecond = 1000 * Microsecond   (1e6)
Second = 1000 * Millisecond         (1e9)
Minute = 60 * Second                (60*1e9)
Hour = 60 * Minute                  (60*60*1e9)
*/
func main() {
    starttime := time.Now()
    //time.After返回一个chan,指定时间后通过chan传回即时时间
    fmt.Println(<-time.After(2*time.Second))
    //time.Sleep 阻塞指定时间
    time.Sleep(2 * time.Second)
    //time.ParseDuration将一个表示时间的字符串解析，每一个字符串是带有单位的十进制数序列，每个数字带有可选的单位或小数位。
    //合法的单位是"ns", "us" , "ms", "s", "m", "h"。返回Duration值
    d, err := time.ParseDuration("1h2m3s4ms5.8us6ns")
    fmt.Println(d, err)
    //返回从时间t到当前时间的间隔，time.Now().Sub(t)的简写
    fmt.Println(time.Now().Sub(starttime))
    fmt.Println(time.Since(starttime))

    //time.Tick 类似time.After,但可以持续使用
    c := time.Tick(1 * time.Second)
    for now := range c {
        fmt.Printf("%v\n", now)
    }
}
