package main
import "time"

//goroutine使用go便捷地实现多协程

func wait(c chan bool, i int) {
    time.Sleep(time.Second*i)
    c <- true
}

func main() {
    var c chan bool
    go wait(c, 5)
    <-c
}