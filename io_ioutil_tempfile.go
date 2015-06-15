package main

import "fmt"
import "io/ioutil"
/*
本函数主要是用来在指定的dir目录下创建一个新的、使用prefix为前缀的临时文件，并以读写模式打开该文件并返回os.File指针。 如果dir是空字符串，TempFile使用默认用于临时文件的目录（参见os.TempDir函数）。 如果多个程序调用该函数的话，将会创建不同的临时文件（因此是线程安全的）。调用本函数的程序有责任在不需要临时文件时摧毁它.
*/
func main() {
    f, e := ioutil.TempFile(".", "temp")
    defer f.Close()
    fmt.Println(f.Name())
    if e != nil {
        fmt.Println("create tempFile error")
        return
    }
    f.WriteString("hello world!")
    b, e1 := ioutil.ReadFile(f.Name())
    if e1 != nil {
        fmt.Println("read error")
        return
    }
    fmt.Println(string(b))
}