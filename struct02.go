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
    person // 匿名字段,实现继承person的所有字段
    school string
}

type teacher struct {
    person // 匿名字段
    age int // 与匿名字段的age属性同名时,teacher.age与teacher.person.age为不同的两个值
}

func main() {
    mark := student{person{"mark", 25}, "Computer Science"}//匿名继承初始化必须使用被继承struct赋值
    fmt.Println(mark)
    fmt.Println(mark.name,mark.age,mark.school)//匿名继承，可直接使用被继承struct的属性
    
    mark.school="high school"
    mark.person=person{"mark", 26}
    mark.person.age += 1
    fmt.Println(mark)
    mark.age+=1
    fmt.Println(mark)

    Lily := teacher{person{"Lily", 20}, 10}
    fmt.Println(Lily)
    fmt.Println(Lily.age)  //当struct的字段与继承的字段重名时,直接调用的将是struct自身的字段
    fmt.Println(Lily.person.age) //而继承的字段由于重名,必须使用完整的继承关系调用

}
