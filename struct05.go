package main
import (
    "fmt"
)
//struct的方法的继承
//struct被继承后,方法同样被继承. 新struct的新方法不能被原struct使用.
type person struct {
    name string
    age  int
}

type student person

//p为指针, 则可传入person, 也可以传入*person, 效果同为*person.所以函数内的操作可以改变原对象
func (p *person)sayname() {
    fmt.Println(p.name)
}

//p为普通对象, 则仅可传入非指针对象的拷贝, 无法修改原对象
func (s student)sayage() {
    fmt.Println(s.age)
}

func main() {
    mark := person{"mark", 25}
    fmt.Println(mark)
    mark.sayname()
    //mark.sayage()  person无sayage方法.
    Lily := student{"Lily", 21}
    fmt.Println(Lily)
    Lily.sayname() //student继承了sayname方法
    Lily.sayage()   //sayage方法为student独有
}