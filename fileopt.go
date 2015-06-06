package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func create() {
	finame := "file_test.txt"
	fi, err := os.Create(finame)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
}

func write() {
	finame := "file_test.txt"
	fi, err := os.OpenFile(finame, os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fi.WriteString("file test!\n")
	fi.Write([]byte("file test byte!\n"))
}

func add() {
	finame := "file_test.txt"
	fi, err := os.OpenFile(finame, os.O_RDWR|os.O_APPEND, 644)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fi.WriteString("add string test!\n")
}

func read() {
	finame := "file_test.txt"
	fi, err := os.Open(finame)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	fmt.Print(string(fd))
}

func del() {
	finame := "file_test.txt"
	err := os.Remove(finame)
	if err != nil {
		fmt.Println("finame", err)
	}

}

func main() {
	create()
	write()
	add()
	read()
	del()
}
