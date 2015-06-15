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
    gobook := Book{"Go语言编程", []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"}, "ituring.com.cn", true,        9.99    }
    b, err := json.Marshal(gobook)
    if err!=nil{
        fmt.Println(err.Error())
    }
    fmt.Println(string(b))
}
