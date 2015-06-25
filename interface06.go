package main
import "fmt"
/*
接口查询
接口查询是否成功，要在运行期才能够确定。他不像接口的赋值，编译器只需要通过静态类型检查即可判断赋值是否可行。
*/

func main() {
    var a interface{}
    a = int8(1)

    if value, ok := a.(int8); ok { //断言a值的类型,可返回该类型的a值.
        fmt.Println("I guessed I am int8, and I am right.")
        fmt.Println("value is ", value)
    }

    switch a.(type){// a.(type) 仅适用于switch下,并且a必须为接口类型
        case int:
        fmt.Println("I am int.")
        case int8:
        fmt.Println("I am int8.")
        case int32:
        fmt.Println("I am int32.")
        case int64:
        fmt.Println("I am int64.")
        default:
        fmt.Println("What am I?")
    }
}
