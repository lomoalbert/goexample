package main

import (
    "fmt"
    "reflect"
)

type MyStruct struct {
    name string
}

func (this *MyStruct)GetName() string {
    return this.name
}



func main() {
    //反射用于获取对象的type和value两个属性
    s := "this is string"
    fmt.Println("reflect.TypeOf(string):", reflect.TypeOf(s))
    fmt.Println("reflect.ValueOf(string)", reflect.ValueOf(s))

    var x float64 = 3.4
    fmt.Println("reflect.TypeOf(float64):", reflect.TypeOf(x))
    fmt.Println("reflect.ValueOf(float64):", reflect.ValueOf(x))
    fmt.Println("kind is float64:", reflect.TypeOf(x).Kind() == reflect.Float64)
    fmt.Println("float64 value:", reflect.ValueOf(x).Float())
    fmt.Println("canset of float64:", reflect.ValueOf(x).CanSet())
    fmt.Println("type of *float64:", reflect.TypeOf(&x))
    fmt.Println("canset of *float64:", reflect.ValueOf(&x).CanSet())
    p := reflect.ValueOf(&x)
    v := p.Elem()
    fmt.Println("canset of elem of *float64:", v.CanSet())
    v.SetFloat(4.3)
    fmt.Println("after setfloat float64 value:", x)


    var i int = 3
    fmt.Println("reflect.TypeOf(int):", reflect.TypeOf(i))
    fmt.Println("reflect.ValueOf(int):", reflect.ValueOf(i))
    fmt.Println("kind is int:", reflect.TypeOf(i).Kind() == reflect.Int)
    fmt.Println("int value:", reflect.ValueOf(i).Int())

    a := new(MyStruct)
    a.name = "Mystruct的name"
    fmt.Println("reflect.TypeOf(MyStruct):", reflect.TypeOf(a))
    fmt.Println("reflect.ValueOf(MyStruct):", reflect.ValueOf(a))

    typ := reflect.TypeOf(a)
    fmt.Println("typ.NumMethod():", typ.NumMethod())
    b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
    fmt.Println(b)

    //静态类型的value的type是原类型
    type Myint int
    var myint Myint = 4
    fmt.Println("type of myint:", reflect.TypeOf(myint))
    fmt.Println("type of value of myint:", reflect.ValueOf(myint).Kind())

}

/*
reflect.TypeOf(string): string
reflect.ValueOf(string) this is string
reflect.TypeOf(float64): float64
reflect.ValueOf(float64): <float64 Value>
kind is float64: true
float64 value: 3.4
canset of float64: false
type of *float64: *float64
canset of *float64: false
canset of elem of *float64: true
after setfloat float64 value: 4.3
reflect.TypeOf(int): int
reflect.ValueOf(int): <int Value>
kind is int: true
int value: 3
reflect.TypeOf(MyStruct): *main.MyStruct
reflect.ValueOf(MyStruct): <*main.MyStruct Value>
typ.NumMethod(): 1
[Mystruct的name]
type of myint: main.Myint
type of value of myint: int
*/