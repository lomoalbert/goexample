package main
import "fmt"
//switch是最灵活的一种控制语句,表达式可以不是常量或者字符串，也可以沒有表达式，如果没有表达式则如同if-else-else。
//一般用法:和其他语言的Switch基本一样，不同的不需要在每个Case中添加Break,而是隐藏了Break,当然你可以显示加入break
//fallthrough 与 break相反, fallthrough将忽略下一个case的判断而直接执行.并且fallthrough必须是当前case的最后一条语句,并且不能放在switch最后一句.
func main() {
    var ch rune
    var cl string
    ch = '0'
    switch ch {
        case '0':
        cl = "0"
        case '1':
        cl = "1"
        break //可以添加,默认自动添加,除非fallthrough
        default:
        cl = "Other Char"
    }
    fmt.Println(cl)

    // switch没有表达式时 与 if-else是相同的
    // switch默认带break效果
    var i int
    i = 3
    switch {
        case i>0:
        fmt.Println("i > 0")
        case i>1:
        fmt.Println("i > 1")
        case i>2:
        fmt.Println("i > 2")
        case i>3:
        fmt.Println("i > 3")
        case i>4:
        fmt.Println("i > 4")
    }


    // fallthrough 忽略下一个case的判断语句 而直接执行
    switch {
        case i>0:
        fmt.Println("i > 0")
        fallthrough
        case i>1:
        fmt.Println("i > 1")
        fallthrough
        case i>2:
        fmt.Println("i > 2")
        fallthrough
        case i>3:
        fmt.Println("i > 3")
        fallthrough
        case i>4:
        fmt.Println("i > 4")
    }

    //如果多个匹配结果所对应的代码段一样，则可以在一个case中并列出所有的匹配项
    switch ch {
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        cl = "Int"
        case 'A', 'B', 'C', 'D', 'a', 'b', 'c', 'd', 'e':
        cl = "ABC"
        default:
        cl = "Other Char"
    }
    fmt.Println(cl)

    //同样switch可以没有表达式，在 Case 中使用布尔表达式，这样形如 if-else
    switch {
        case '0' <= ch && ch <= '9':
        cl = "Int"
        case ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z'):
        cl = "ABC"
        default:
        cl = "Other Char"
    }
    fmt.Println(cl)
}