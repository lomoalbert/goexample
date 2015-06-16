package main

import (
    "io"
    "fmt"
)

type user string

func (s user) Read(b []byte) (int, error) {
    t := []byte(s)
    i := 0
    for ; i<len(t) && i<len(b); i++ {
        b[i]=t[i]
    }
    return i, nil
}
func main() {
    var s user = "abcdefg"
    b := make([]byte, 4)
    n, err := io.ReadFull(s, b)
    fmt.Println(n, err, string(b))
}