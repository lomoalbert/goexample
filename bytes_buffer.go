package main
import (
    "bytes"
    "fmt"
    "os"
)
func main() {
    bb := bytes.NewBuffer([]byte("Hello World!"))
    b := make([]byte, 32)

    bb.Read(b)
    fmt.Printf("%s\n", b) // Hello World!

    bb.WriteString("New Data!\n")
    bb.WriteTo(os.Stdout) // New Data!

    bb.WriteString("Third Data!")
    bb.ReadByte()
    fmt.Println(bb.String()) // hird Data!
    bb.UnreadByte()
    fmt.Println(bb.String()) // Third Data!
}