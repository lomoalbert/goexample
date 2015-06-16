package main
import (
    "bytes"
    "fmt"
    "os"
    "io"
)
func main() {
    bb := bytes.NewBuffer([]byte("Hello World!"))
    b := make([]byte, 32)

    //bb.Read(b)
    n, err := io.ReadFull(bb, b)
    fmt.Println(n, err)
    fmt.Printf("%s\n", b) // Hello World!

    bb.WriteString("New Data!\n")
    bb.WriteTo(os.Stdout) // New Data!

    bb.WriteString("Third Data!")
    bb.ReadByte()
    fmt.Println(bb.String()) // hird Data!
    bb.UnreadByte()
    fmt.Println(bb.String()) // Third Data!
}