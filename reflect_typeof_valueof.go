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
    s := "this is string"
    fmt.Println("reflect.TypeOf(string):", reflect.TypeOf(s))

    fmt.Println("reflect.ValueOf(string)", reflect.ValueOf(s))
    var x float64 = 3.4
    fmt.Println("reflect.TypeOf(float64):", reflect.TypeOf(x))
    fmt.Println("reflect.ValueOf(float64):", reflect.ValueOf(x))
    fmt.Println("kind is float64:", reflect.TypeOf(x).Kind() == reflect.Float64)
    fmt.Println("float64 value:", reflect.ValueOf(x).Float())

    a := new(MyStruct)
    a.name = "Mystruct的name"
    fmt.Println("reflect.TypeOf(MyStruct):", reflect.TypeOf(a))
    fmt.Println("reflect.ValueOf(MyStruct):", reflect.ValueOf(a))

    typ := reflect.TypeOf(a)
    fmt.Println("typ.NumMethod():", typ.NumMethod())
    b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
    fmt.Println(b)

}

/*
/tmp/reflect_typeof_valueof.go0go
string
-------------------
this is string
<float64 Value>
-------------------
1
-------------------
yejianfeng

*/