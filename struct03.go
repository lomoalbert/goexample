package main
import (
    "fmt"
)
//struct的方法
//Reciver 默认以值传递，而非引用传递，还可以是指针
//func (r ReciverType) funcName(params) (results) { ... }
//指针作为Receiver会对实例对象的内容发生操作，而普通类型作为Receiver仅仅是以副本作为操作对象，而不对原实例对象发生操作
/*
接口去调用结构体的方法时需要针对接受者的不同去区分，即：

接收者是指针*T时，接口实例必须是指针
接收者是值 T时，接口实力可以是指针也可以是值
接口的定义和类型转换与接收者的定义是关联的

*/

type person struct {
    name string
    age  int
}

//普通类型作为Receiver仅仅是以副本作为操作对象，而不对原实例对象发生操作
func (p person) growup01() {
    p.age+=1
    fmt.Println(p)
}

//如果一个method的receiver是*T，调用时，可以传递一个T类型的实例变量V，而不必用&V去调用这个method
func (p *person) growup02() {
    p.age+=1
    fmt.Println(p)
}

func main() {
    mark := person{"mark", 25}
    mark.growup01()     //{mark 26}
    fmt.Println(mark)   //{mark 25}
    (&mark).growup01()  //{mark 26}
    fmt.Println(mark)   //{mark 25}
    mark.growup02()     //&{mark 26}
    fmt.Println(mark)   //{mark 26}
    (&mark).growup02()     //&{mark 26}
    fmt.Println(mark)   //{mark 26}
}