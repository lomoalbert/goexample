package main
import "fmt"
/*
接口的赋值
接口赋值在Go语言中分为两种情况: 1.将对象实例赋值给接口 2.将一个接口赋值给另外一个接口


将对象实例赋值给接口
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

/*
go语言可以根据下面的函数:
func (a Integer) Less(b Integer) bool

自动生成一个新的Less()方法
func (a *Integer) Less(b Integer) bool

这样，类型*Integer就既存在Less()方法，也存在Add()方法，满足LessAdder接口。 而根据
func (a *Integer) Add(b Integer)】

这个函数无法生成以下成员方法：
func(a Integer) Add(b Integer) {
（&a).Add(b)
}

因为(&a).Add()改变的只是函数参数a,对外部实际要操作的对象并无影响(值传递)，这不符合用户的预期。所以Go语言不会自动为其生成该函数。因此类型Integer只存在Less()方法，缺少Add()方法，不满足LessAddr接口。（可以这样去理解：指针类型的对象函数是可读可写的，非指针类型的对象函数是只读的）
*/