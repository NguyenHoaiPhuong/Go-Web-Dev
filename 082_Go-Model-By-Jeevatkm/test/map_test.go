package test

import (
	"GO-WEB-DEV/082_Go-Model-By-Jeevatkm/product"
	"fmt"
	"testing"

	model "gopkg.in/jeevatkm/go-model.v1"
)

func TestMapProductA(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestMapProductA")
	A := &product.A{
		ID:   "1234",
		Name: "Hair Dry",
	}
	srchResMap, err := model.Map(A)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Printf("Search Result Map: %#v\n", srchResMap)
	fmt.Printf("ID: %v\n", srchResMap["ID"])
	fmt.Printf("Name: %v\n", srchResMap["Name"])
}
