package main
import "fmt"
/*
接口的赋值
接口赋值在Go语言中分为两种情况: 1.将对象实例赋值给接口 2.将一个接口赋值给另外一个接口

将一个接口赋值给另外一个接口
在Go语言中，只要两个接口拥有相同的方法列表(次序不同不要紧),那么它们就等同的，可以相互赋值。 如果A接口的方法列表时接口B的方法列表的子集，那么接口B可以赋值给接口A，但是反过来是不行的，无法通过编译。
*/

type LesssAdder interface {
    Less(b Integer) bool
    Add(b Integer)
}

type Integer int

func (a Integer) Less(b Integer) bool {
    return a < b
}

func (a *Integer) Add(b Integer) {
    *a += b
}

func main() {

    var a Integer = 1
    var b LesssAdder = &a
    fmt.Println(b)

    //var c LesssAdder = a
    //Error:Integer does not implement LesssAdder
    //(Add method has pointer receiver)
}
