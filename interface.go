package main
import "fmt"
import (
    "io"
)
// todo: interface的用途描述
/*
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
*/

type Myinterface struct {
    readstring string
    writestring string
}

func (mi Myinterface) Read(p []byte) (int, error) {
    p = []byte(mi.readstring)
    return len(p), nil
}

func (mi Myinterface) Write(p []byte) (n int, err error) {
    mi.writestring = string(p)
    return len(p), nil
}
func main() {


    mi := Myinterface{readstring:"readingstring", writestring : "writingstring"}
    b := make([]byte, 10)
    io.ReadFull(mi, b)
    fmt.Println(b)
}