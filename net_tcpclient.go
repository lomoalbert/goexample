package main

//客户机
import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
    "strings"
)

//连接服务器
func main() {
    //接通
    conn, err := net.Dial("tcp", "localhost:7777")
    checkError(err)
    fmt.Println("连接成功！\n")
    //输入
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("你是谁？")
    name, _ := inputReader.ReadString('\n')
    //
    trimName := strings.Trim(name, "\r\n")
    conn.Write([]byte(trimName + " 接入了\n "))
    fmt.Println("我们来聊天吧！按quit退出")
    for {
        //读一行
        input, _ := inputReader.ReadString('\n')
        trimInput := strings.Trim(input, "\r\n")
        //如果quit就退出
        if trimInput == "quit" {
            fmt.Println("再见")
            conn.Write([]byte(trimName + " 退出了 "))
            return
        }
        //写出来
        _, err = conn.Write([]byte(trimName + " says " + trimInput))
    }
}

//检查错误
func checkError(err error) {
    if err != nil {
        log.Fatal("an error!", err.Error())
    }
}
