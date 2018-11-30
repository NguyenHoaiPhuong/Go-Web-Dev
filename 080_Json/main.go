package main

import (
	"GO-WEB-DEV/080_Json/jsonfunc"
	"GO-WEB-DEV/080_Json/model"
	"bytes"
	"fmt"
)

func main() {
	akagi := &model.Person{}
	err := jsonfunc.ReadFromFile("test/person.json", akagi)
	if err != nil {
		fmt.Println(err.Error())
	}

	bs, err := jsonfunc.ConvertToJSON(akagi)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = jsonfunc.WriteToFile(bytes.NewReader(bs), "test/converted_person.json")
	if err != nil {
		fmt.Println(err.Error())
	}
}
