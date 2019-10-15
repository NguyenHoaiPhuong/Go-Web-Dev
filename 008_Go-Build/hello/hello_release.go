// +build !debug

package hello

import "fmt"

func Hello() {
	fmt.Println("Release build")
}
