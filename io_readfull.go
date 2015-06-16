package main

import (
    "io"
    "fmt"
    "errors"
)

type user string

func (s user) Read(b []byte) (int, error) {
    t := []byte(s)
    i := 0
    for ; i<len(t) && i<len(b); i++ {
        b[i]=t[i]
    }
    switch i {
        case len(b):
        return i, nil
        case len(t):
        return i, io.EOF
        default:
        return i, errors.New("Read Fail")
    }
}
func main() {
    var s user = "abcdefg"
    b := make([]byte, 10)
    n, err := io.ReadFull(s, b)
    fmt.Println(n, err, string(b))
}