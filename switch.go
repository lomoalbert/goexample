package goexample
import "fmt"
//switch是最灵活的一种控制语句,表达式可以不是常量或者字符串，也可以沒有表达式，如果没有表达式则如同if-else-else。
//一般用法:和其他语言的Switch基本一样，不同的不需要在每个Case中添加Break,而是隐藏了Break,当然你可以显示加入break

func main() {
    var ch rune
    var cl string
    ch = '0'
    switch ch {
        case '0':
        cl = "0"
        case '1':
        cl = "1"
        break //可以添加
        default:
        cl = "Other Char"
    }
    fmt.Println(cl)

    switch {
        case ch=='0':
        cl="0"
        case ch=='1':
        cl="1"
        default:
        cl="Other Char"
    }
    fmt.Println(cl)


}