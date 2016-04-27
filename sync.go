package main

import (
    "fmt"
    "sync"
)

func main() {
    wg := new(sync.WaitGroup)
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go done(wg)
    }
    wg.Wait()
    fmt.Println("exit")
}

func done(wg sync.WaitGroup) {
    //dosomething
    wg.Done()
}
