package main
import (
    "fmt"
)
//struct的方法
//struct不仅可以使用struct作为匿名字段，自定义类型、内置类型都可以作为匿名字段,而且可以在相应字段上做函数操作
type person struct {
    name string
    age  int
}

func (p person) growup01() {
    p.age+=1
    fmt.Println(p)
}

func (p *person) growup02() {
    p.age+=1
    fmt.Println(p)
}

func main() {
    mark := person{"mark", 25}
    mark.growup01()
    fmt.Println(mark)
}