package main

import (
    "time"
    "fmt"
    "strings"
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
    p := fmt.Println
    starttime := time.Now()
    //time.After返回一个chan,指定时间后通过chan传回即时时间
    fmt.Println(<-time.After(time.Second))
    //time.Sleep 阻塞指定时间
    time.Sleep(time.Second)
    //time.ParseDuration将一个表示时间的字符串解析，每一个字符串是带有单位的十进制数序列，每个数字带有可选的单位或小数位。
    //合法的单位是"ns", "us" , "ms", "s", "m", "h"。返回Duration值
    d, err := time.ParseDuration("1h2m3s4ms5.8us6ns")
    fmt.Println(d, err)
    
    t := "2015-06-15T14:45:00Z"
    fmt.Println(time.RFC3339)
    ti , _ := time.Parse(time.RFC3339,t)
    fmt.Println(ti)
    //Duration值默认单位为nanoseconds,但也可以传唤为其他单位,以及字符串
    fmt.Println(d.Hours(), d.Minutes(), d.Seconds(), d.Nanoseconds(), d.String())
    //返回从时间t到当前时间的间隔，time.Now().Sub(t)的简写
    fmt.Println(time.Now().Sub(starttime))
    fmt.Println(time.Since(starttime))
    //返回时间的时区信息
    fmt.Println(starttime.Location())

    fmt.Println(strings.Repeat("=", 10))

    now := time.Now()
    p(now)
    p(now.Zone())

    d = time.Duration(7200 * 1000 * 1000 * 1000)
    p(d)
    then := time.Date(2013, 1, 7, 20, 34, 58, 651387237, time.UTC)

    p(then)
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())
    p(then.Weekday())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    p(then.Date())
    p(then.ISOWeek())
    p("----------")
    p(now.UTC())
    p(now.Local())
    p(now.Location())
    p(now.Zone())
    p(now.Unix())
    p(time.Unix(now.Unix(), 0))
    p(now.UnixNano())
    p(time.Unix(0, now.UnixNano()))
    p(now.GobEncode())
    p(now.MarshalJSON())
    p(time.Since(now))
    p("----------")
    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())
    p(then.Add(diff))
    p(then.Add(-diff))

    p(d)
    p(d.Hours())
    p(d.Minutes())
    p(d.Seconds())
    p(d.Nanoseconds())
    p(then.Add(d))

    //time.Tick 类似time.After,但可以持续使用
    c := time.Tick(1 * time.Second)
    for now := range c {
        fmt.Printf("%v\n", now)
    }
}
