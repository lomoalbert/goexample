package main
import "time"

//goroutine使用go便捷地实现多协程

func wait(c chan bool, t time.Duration) {
    time.Sleep(t)
    c <- true
}

func main() {
    c := make(chan bool)
    go wait(c, 5*time.Second)
    <-c

}