package main
import (
    "fmt"
)
//struct的匿名字段 即继承
//struct不仅可以使用struct作为匿名字段，自定义类型、内置类型都可以作为匿名字段,而且可以在相应字段上做函数操作
type person struct {
    name string
    age  int
}

type student struct {
    person // 匿名函数,实现继承person结构
    school string
}

func main() {
    mark := student{person{"mark", 25}, "Computer Science"}//匿名继承初始化必须使用被继承struct
    fmt.Println(mark)
    fmt.Println(mark.name,mark.age,mark.school)//匿名继承，可直接使用被继承struct的属性
    
    mark.school="high school"
    mark.person=person{"mark", 26}
    mark.person.age += 1
    fmt.Println(mark)
    mark.age+=1
    fmt.Println(mark)



}
