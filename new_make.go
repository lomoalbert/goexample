package main
import (
    "fmt"
    "reflect"
)
//new 的作用是初始化一个指向类型的指针(*T)，make 的作用是为 slice，map 或 chan 初始化并返回引用(T)。

//func new(Type) *Type
//内建函数 new 用来分配内存，它的第一个参数是一个类型，不是一个值，它的返回值是一个指向新分配类型零值的指针

//func make(Type, size IntegerType) Type
/*
内建函数 make 用来为 slice，map 或 chan 类型分配内存和初始化一个对象(注意：只能用在这三种类型上)，跟 new 类似，第一个参数也是一个类型而不是一个值，跟 new 不同的是，make 返回类型的引用而不是指针，而返回值也依赖于具体传入的类型，具体说明如下：
    Slice: 第二个参数 size 指定了它的长度，它的容量和长度相同。
    你可以传入第三个参数来指定不同的容量值，但必须不能比长度值小。
    比如 make([]int, 0, 10)
    Map: 根据 size 大小来初始化分配内存，不过分配后的 map 长度为 0，如果 size 被忽略了，那么会在初始化分配内存时分配一个小尺寸的内存
    Channel: 管道缓冲区依据缓冲区容量被初始化。如果容量为 0 或者忽略容量，管道是没有缓冲区的.没有缓冲区时,输入输出要相互等待,必须同时进行;有缓冲区时,输入仅在缓冲满,输出仅在缓冲空的时候等待.

*/

func main() {
    someInt := 1
    otherInt := new(int)
    fmt.Println(someInt, reflect.TypeOf(someInt))
    fmt.Println(otherInt, reflect.TypeOf(otherInt))

    someSlice := []int{0, 0, 0, 0, 0}
    otherSlice := make([]int, 5)
    fmt.Println(someSlice, reflect.TypeOf(someSlice))
    fmt.Println(otherSlice, reflect.TypeOf(otherSlice))


}
