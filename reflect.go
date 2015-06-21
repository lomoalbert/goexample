package main
import "fmt"
import "reflect"
func hello() {
    fmt.Println("hello world!")
}

func main() {
    hl := hello
    fv := reflect.ValueOf(hl)
    fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
    fv.Call(nil)
}