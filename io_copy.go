package main
import (
    "io"
    "fmt"
    "os"
)

func main() {
    srcFile, _ := os.Open("copySrc.txt")
    destFile, _ := os.Create("copyDest.txt")
    written, err := io.Copy(destFile, srcFile)
    if err == nil {
        fmt.Println("Copy Success! total", written, "bytes")
    }
}