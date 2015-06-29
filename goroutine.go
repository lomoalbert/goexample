package main
import "time"

//goroutine使用go便捷地实现多协程

func wait(c chan bool, t time.Duration) {
    time.Sleep(t)
    c <- true
}

func main() {
    c := make(chan bool) //无缓冲区,输入输出必须同时执行
    go wait(c, 5*time.Second)
    <-c

}