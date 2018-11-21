package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	msg := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present.\n"

	strReader := strings.NewReader(msg)
	fmt.Printf("strReader: %p %p\n", strReader, &strReader)
	_, err := io.Copy(os.Stdout, strReader)
	if err != nil {
		fmt.Println(err)
		return
	}

	bReader := bytes.NewReader([]byte(msg))
	fmt.Printf("bReader: %p %p\n", bReader, &bReader)
	//bBuffer := bytes.NewBuffer([]byte(msg))
	_, err = io.Copy(os.Stdout, bReader)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.Get("https://www.google.com.vn")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
