package main

import (
	"fmt"
)

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func spec(v interface{}) {
	fmt.Println(v)
}

func main() {
	kiki := dog{
		animal{sound: "gogo"},
		true,
	}
	mino := cat{
		animal{sound: "meow"},
		true,
	}
	spec(kiki)
	spec(mino)
}
