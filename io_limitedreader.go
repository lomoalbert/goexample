package main;

import (
    "io"
    "fmt"
    "os"
)

func main() {
    reader, _ := os.Open("readFile.txt")
    limitReader := io.LimitReader(reader, 20)
    var n, total int
    var err error
    p := make([]byte, 15)
    for {
        n, err = limitReader.Read(p)
        if err == io.EOF {
            fmt.Println("Read end total", total)
            break
        }
        total = total + n
        fmt.Println("Read value:", string(p[0:n]))
        fmt.Println("Read count:", n)
    }
}