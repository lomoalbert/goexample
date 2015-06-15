package main

import "fmt"
import "io/ioutil"
import "strings"

func main() {
    s := strings.NewReader("hello world!")
    b, e := ioutil.ReadAll(s)
    if e != nil {
        fmt.Printf("%v\n", e)
        return
    }
    fmt.Println(string(b))
}