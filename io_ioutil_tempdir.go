package main

import "fmt"
import "io/ioutil"
/*
本函数主要是用来在指定的dir目录里创建一个新的、使用prfix作为前缀的临时文件夹，并返回文件夹的路径。 如果dir是空字符串，TempDir使用默认用于临时文件的目录（参见os.TempDir函数）。 如果多个程序调用该函数的话，将会创建不同的临时目录（因此是线程安全的）。调用本函数的程序有责任在不需要临时文件夹时摧毁它。
*/
func main() {
    f, e := ioutil.TempDir("d:/goTest", "temp")
    if e != nil {
        fmt.Println("create tempDir error")
        return
    }
    fmt.Println(f)
}