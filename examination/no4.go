package main

import "fmt"
import "time"

/*程序 04：这一年的第几天

题目：输入某年某月某日，判断这一天是这一年的第几天？
1.程序分析：以3月5日为例，应该先把前两个月的加起来，然后再加上5天即本年的第几天，特殊情况，闰年且输入月份大于3时需考虑多加一天。
2.程序源代码：*/


func main () {
    var date string
    fmt.Println("输入日期,如20150130:")
    fmt.Scan(&date)
    today,_ := time.Parse("20060102",date)
    firstday,_ := time.Parse("2006",date[:4])
    fmt.Printf("是今年的第%d天!\n",today.Sub(firstday)/time.Hour/24+1)
}