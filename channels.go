package main

import "fmt"

func Count(ch, counter chan int) {
    select {
    case ch <- 0:
    case ch <- 1:
    }
}

func main() {
    ch := make(chan int, 100)
    counter := make(chan int, 100)
    for i := 0; i < 100; i++ {
        go Count(ch, counter)
    }
    for i := 0; i < 100; i++ {
        fmt.Println(<-ch)
    }
}