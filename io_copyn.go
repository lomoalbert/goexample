package main

import (
    "io"
    "fmt"
    "os"
)

func main() {
    srcFile, _ := os.Open("copySrc.txt")
    destFile, _ := os.Create("copyDest.txt")
    written, err := io.CopyN(destFile, srcFile, 15)
    if err == nil {
        fmt.Println("Copy Success! total", written, "bytes")
    }
    if err == io.EOF {
        fmt.Println("Copy all total", written, "bytes")
    }
}