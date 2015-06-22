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


    type Html []interface{}

    html := make(Html, 5)
    html[0] = "div"
    html[1] = "span"
    html[2] = []byte("script")
    html[3] = "style"
    html[4] = "head"
    //断言方式
    for index, element := range html {
        if value, ok := element.(string); ok {
            fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
        } else if value, ok := element.([]byte); ok {
            fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
        }
    }
    //switch方式
    for index, element := range html {
        switch value := element.(type) {
            case string:
            fmt.Printf("html[%d] is a string and its value is %s\n", index, value)
            case []byte:
            fmt.Printf("html[%d] is a []byte and its value is %s\n", index, string(value))
            case int:
            fmt.Printf("invalid type\n")
            default:
            fmt.Printf("unknown type\n")
        }
    }
}