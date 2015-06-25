package main
import "fmt"
/*
接口的元素是方法,既可以有0个,也可以有多个
实现了接口指定的方法也就实现了接口,就可以代入到以该接口为参数的运算中.
包含0个方法的接口为空接口.任何对象均实现空接口,空接口对象可以接收任意类型的值,
未被初始化的 interface 类型变量的零值为 nil
*/

func main() {
    var empty interface{}
    empty = 10
    fmt.Printf("now EmptyInterface is of type %T\n", empty)
    empty = "string"
    fmt.Printf("now EmptyInterface is of type %T\n", empty)
    empty = 'a'
    fmt.Printf("now EmptyInterface is of type %T\n", empty)
    empty = []byte("byte")
    fmt.Printf("now EmptyInterface is of type %T\n", empty)
}