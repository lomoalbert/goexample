package main
import (
    "fmt"
)
//struct方法的继承
type person struct {
    name string
    age  int
}

func (p person) sayname() {
    fmt.Println(p.name)
}

type student struct {
    person //studen继承person的同时也继承了person的方法
    school string
}

func main() {
    mark := person{"mark", 25}
    mark.sayname()
    mark_student := student{mark, "Computer Science"}
    mark_student.sayname()
    mark_student.person.sayname()
}