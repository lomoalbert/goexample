package main
import "fmt"

func main() {
    var val interface{} = "good"
    //Comma-ok断言的语法是：value, ok := element.(T)。element必须是接口类型的变量，T是普通类型。如果断言失败，ok为false，否则ok为true并且value为变量的值。
    value, ok := val.(string)
    if ok {

        fmt.Println(value)
    }
    //其实Comma-ok断言还支持另一种简化使用的方式：value := element.(T)。但这种方式不建议使用，因为一旦element.(T)断言失败，则会产生运行时错误。
    value = val.(string)
    fmt.Println(value)
    //fmt.Println(val.(int))   //将出现程序错误
}