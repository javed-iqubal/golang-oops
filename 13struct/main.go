package main

import "fmt"

// interface embeding- Golang does not support inheritence, composition implement interface

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReaderWriter interface {
	Reader
	Writer
}

type File struct {
	Name string
}

func (f File) Read() {
	fmt.Println("Reading from file ", f.Name)
	//return "Read"
}

func (f File) Write() {
	fmt.Println("Writing into file ", f.Name)
	//return "Write"
}

func Process(rw ReaderWriter) {
	rw.Read()
	rw.Write()
}

func main() {

	// file := File{Name: "data.txt"}

	var rw ReaderWriter = File{Name: "test"}

	rw.Read()
	rw.Write()
	Process(rw)

}
