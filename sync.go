package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        add(wg)
    }

    for i := 0; i < 100; i++ {
        go done(wg)
    }
    wg.Wait()
    fmt.Println("exit")
}

func add(wg sync.WaitGroup) {
    wg.Add(1)
}

func done(wg sync.WaitGroup) {
    wg.Done()
}
