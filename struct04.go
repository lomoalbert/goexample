package main
import (
    "fmt"
)
//struct方法的继承
type person struct {
    name string
    age  int
}

//struct
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
    mark.growup02()     //&{mark 26}
    fmt.Println(mark)   //{mark 26}
}