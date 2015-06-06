package main
import "fmt"
func MyPrintf(args ...interface{}) {
    for _, arg := range args {
        fmt.Println(type(arg))
        switch arg.(type) {
            case int:
            fmt.Println(arg, "is an int value.")
            case string:
            fmt.Println(arg, "is an string value.")
            case int64:
            fmt.Println(arg, "is an int64 value.")
            default:
            fmt.Println(arg, "is an unkown value.")
        }
    }
}

func main() {
    var v1 int = 1
    var v2 int64 = 234
    var v3 string = "hello"
    var v4 float32 = 1.234
    MyPrintf(v1, v2, v3,    v4)
}