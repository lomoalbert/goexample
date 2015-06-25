package main
import "fmt"
/*
接口的赋值
*/

type emptyinsterface interface {
}

func main() {
    var v emptyinsterface
    // var v interface{}
    v = 10
    fmt.Printf("now EmptyInterface is of type %T\n", v)
    v = "string"
    fmt.Printf("now EmptyInterface is of type %T\n", v)
    v = 'a'
    fmt.Printf("now EmptyInterface is of type %T\n", v)
    v = []byte("byte")
    fmt.Printf("now EmptyInterface is of type %T\n", v)
}