package main

//服务器端
import (
    "fmt"
    "log"
    "net" //支持通讯的包
)

//开始服务器
func main() {
    //连接主机、端口，采用ｔｃｐ方式通信，监听７７７７端口
    listener, err := net.Listen("tcp", "localhost:7777")
    checkError(err)
    fmt.Println("建立成功!")
    for {
        //等待客户端接入
        conn, err := listener.Accept()
        checkError(err)
        //开一个goroutines处理客户端消息，这是golang的特色，实现并发就只go一下就好
        go doServerStuff(conn)
    }
}

//处理客户端消息
func doServerStuff(conn net.Conn) {
    nameInfo := make([]byte, 512) //生成一个缓存数组
    _, err := conn.Read(nameInfo)
    checkError(err)

    for {
        buf := make([]byte, 512)
        _, err := conn.Read(buf) //读取客户机发的消息
        flag := checkError(err)
        if flag == 0 {
            break
        }
        fmt.Println(string(buf)) //打印出来
    }
}

//检查错误
func checkError(err error) int {
    if err != nil {
        if err.Error() == "EOF" {
            //fmt.Println("用户退出了")
            return 0
        }
        log.Fatal("an error!", err.Error())
        return -1
    }
    return 1
}
