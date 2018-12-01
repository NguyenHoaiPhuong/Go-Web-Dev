package main

import (
	"GO-WEB-DEV/080_Json/jsonfunc"
	"GO-WEB-DEV/098_AutoMapper/model"
	"bytes"
	"errors"
	"fmt"

	automapper "github.com/PeteProgrammer/go-automapper"
)

type personename string

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("Recover: %v\n", r)
		}
	}()

	akagi := &model.Person{}
	err := jsonfunc.ReadFromFile("testfile/person.json", akagi)
	if err != nil {
		err = errors.New("Read From Json File Failed")
		fmt.Println(err.Error())
	}

	aka := &model.Human{}
	//automapper.Map(akagi, aka)
	automapper.MapLoose(akagi, aka)
	bs, err := jsonfunc.ConvertToJSON(aka)
	if err != nil {
		err = errors.New("Convert To Json Byte Slice Failed")
		fmt.Println(err.Error())
	}
	err = jsonfunc.WriteToFile(bytes.NewReader(bs), "testfile/human_converted.json")
	if err != nil {
		err = errors.New("Write To Json File Failed")
		fmt.Println(err.Error())
	}

	akagiFull := &model.PersonFull{}
	//automapper.Map(akagi, akagiFull) // For testing
	automapper.MapLoose(akagi, akagiFull)
	bs, err = jsonfunc.ConvertToJSON(akagiFull)
	if err != nil {
		err = errors.New("Convert To Json Byte Slice Failed")
		fmt.Println(err.Error())
	}
	err = jsonfunc.WriteToFile(bytes.NewReader(bs), "testfile/personful_converted.json")
	if err != nil {
		err = errors.New("Write To Json File Failed")
		fmt.Println(err.Error())
	}

	fmt.Printf("Converting from Person to Human:\n")
	fmt.Printf("Person:\n")
	fmt.Printf("%v\n", *akagi)
	fmt.Printf("Human:\n")
	fmt.Printf("%v\n", *aka)

	fmt.Printf("Converting from Person to PersonFull:\n")
	fmt.Printf("Person:\n")
	fmt.Printf("%v\n", *akagi)
	fmt.Printf("PersonFull:\n")
	fmt.Printf("%v\n", *akagiFull)

	msg := "Nguyen Hoai Phuong"
	var name personename
	automapper.Map(&msg, &name)
	fmt.Println("name:", name)
}
