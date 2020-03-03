package main

import (
	"encoding/json"
	"fmt"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/000_Basic/11_Package/errors"
)

// S struct
type S struct {
	Int       int
	String    string
	ByteSlice []byte
}

func testErrorsPkg() {
	e := errors.New("backend", int32(12), "error message")
	fmt.Println(e)
}

func testStructS() {
	s := &S{42, "Hello World!", []byte{0, 1, 2, 3, 4}}
	out, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	a := new(S)
	err = json.Unmarshal(out, a)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func main() {
	// testErrorsPkg()
	testStructS()
}
