package main
import "fmt"

func main(){
    for range [100]int{} {
        fmt.Println("a")
    }
    for k,v :=range [100]int{} {
        fmt.Println(k,v)
    }

    list:=make([]int,100,200)
    list1:=list[:10]
    fmt.Println(len(list1),cap(list1))
    for k,v :=range list1 {
        fmt.Println(k,v)
    }
}