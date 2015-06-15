package main

import "fmt"
import "encoding/json"

type Book struct {
    Title string
    Authors []string
    Publisher string
    IsPublished bool
    Price float32
}
func main() {
    var book Book
    jsonstring:=[]byte(`{"Title":"Go语言编程","Authors":["XuShiwei","HughLv","Pandaman","GuaguaSong","HanTuo","BertYuan","XuDaoli"],"Publisher":"ituring.com.cn","IsPublished":true,"Price":9.99}`)
    err := json.Unmarshal(jsonstring,&book)
    if err!=nil{
        fmt.Println(err.Error())
    }
    fmt.Println(book)
}
