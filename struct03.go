package main
import (
    "fmt"
)
//struct的方法
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
    mark.growup01() //{mark 26}
    fmt.Println(mark)//{mark 25}
    mark.growup02()//&{mark 26}
    fmt.Println(mark)//{mark 26}
}