package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ReadFromTerminal : read string from stdio terminal
func ReadFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	return text
}

func main() {
	for {
		strChars := ReadFromTerminal()
		str := RandomString(4, strChars)
		fmt.Println(str)
	}
}
