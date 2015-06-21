package main
import "fmt"
import (
    "reflect"
    "strconv"
)
func hello() {
    fmt.Println("hello world!")
}

func prints(i int) string {
    fmt.Println("i =", i)
    return strconv.Itoa(i)
}

func main() {
    fv := reflect.ValueOf(prints)
    params := make([]reflect.Value, 1)  //参数
    params[0] = reflect.ValueOf(20)   //参数设置为20
    rs := fv.Call(params)              //rs作为结果接受函数的返回值
    fmt.Println("result:", rs[0].Interface().(string)) //当然也可以直接是rs[0].Interface()

    hl := hello
    fv1 := reflect.ValueOf(hl)
    fmt.Println("fv1 is reflect.Func ?", fv1.Kind() == reflect.Func)
    fv1.Call(nil)
}