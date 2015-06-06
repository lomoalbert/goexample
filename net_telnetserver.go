package main

//服务器端
import (
    "fmt"
    "log"
    "net" //支持通讯的包
    "os"
    "bufio"
)

type message struct {
    name []byte
    msg []byte
}
//开始服务器
func main() {
    //连接主机、端口，采用ｔｃｐ方式通信，监听23端口
    listener, err := net.Listen("tcp", ":7777")
    check(err)
    fmt.Println("建立成功!")
    msgch := make(chan message)
    connch := make(chan net.Conn)
    go func() {
        for {
            msg := make([]byte, 512)
            mainmsg := message{name:[]byte("admin: "), msg:msg}
            os.Stdin.Read(mainmsg.msg)
            msgch <- mainmsg
        }
    }()
    go sending(msgch, connch)
    for {
        //等待客户端接入
        conn, err := listener.Accept()
        check(err)
        //开一个goroutines处理客户端消息，这是golang的特色，实现并发就只go一下就好
        connch <- conn
        go doclient(msgch, conn,connch)
    }
}

func sending(mesch chan message, connch chan net.Conn) {
    connlist := make([]net.Conn, 0)
    for {
        select {
        case msg := <-mesch :
            for _, c := range connlist {
                _, err := c.Write(append(msg.name, msg.msg...))
                check(err)
            }
        case conn := <-connch:
            flag:=false
            index:=0
            for i,co:=range connlist{
                if conn==co{
                    index=i
                    flag=true
                    break
                }
            }
            if flag{
                connlist=append(connlist[:index],connlist[index+1:]...)
                flag=false
            }else{
                connlist=append(connlist,conn)
            }
        }
    }
}

//处理客户端消息
func doclient(msgch chan message, conn net.Conn,connch chan net.Conn) {
    conn.Write([]byte("What's your name?\r\n"))
    client_reader:=bufio.NewReader(conn)
    nameInfo,err:=client_reader.ReadBytes('\n')
    namelen:=len(nameInfo)
    check(err)
    nameInfo=append(nameInfo[:namelen-2], []byte(": ")...)
    fmt.Println(string(nameInfo), "已连接.")
    conn.Write([]byte("What do you want to say?\r\n"))
    for {
        buf,err:=client_reader.ReadBytes('\n')
        flag := check(err)
        if flag == 0 {
            fmt.Println(string(nameInfo), "退出了.")
            connch <- conn
            break
        }
        msgch <- message{name:nameInfo, msg:buf}
        fmt.Print(string(nameInfo), string(buf)) //打印出来
    }
}

//检查错误
func check(err error) int {
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
