package main
// 接口的标准用法
import "fmt"

type S struct {
    i int
}
func (p S) Get() int {
    return p.i
}
func (p S) Put(v int) {
    p.i = v
}
type I interface{
    Get() int
    Put(int)
}
func f(p I) {
    fmt.Println(p.Get())
    p.Put(1)
}
func main() {
    var s S
    f(s)
    fmt.Println(s.Get())

}