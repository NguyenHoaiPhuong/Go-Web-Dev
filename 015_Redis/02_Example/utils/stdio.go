package utils

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFromTerminal : read string from stdio terminal
func ReadFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	return text
}