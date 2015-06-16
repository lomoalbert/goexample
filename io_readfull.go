package main

import (
    "io"
    "math"
    "fmt"
)

type user string

func (s user) Read(b []byte) (int, error) {
    t := []byte(s)
    l := int(math.Min(float64(len(t)), float64(len(b))))
    for i := 0; i<l; i++ {
        b[i]=t[i]
    }
    return l, nil
}
func main() {
    var s user = "abcdefg"
    b := make([]byte, 4)
    io.ReadFull(s, b)
    fmt.Println(string(b))
}