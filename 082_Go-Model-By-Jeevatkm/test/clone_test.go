package test

import (
	"GO-WEB-DEV/082_Go-Model-By-Jeevatkm/product"
	"fmt"
	"testing"

	model "gopkg.in/jeevatkm/go-model.v1"
)

func TestCloneProductA(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestCloneProductA")
	A := &product.A{
		ID:   "1234",
		Name: "Hair Dry",
	}
	cloneObj, err := model.Clone(A)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	fmt.Printf("Origin Object: %v\n", A)
	fmt.Printf("Clone Object: %v\n", cloneObj)
	fmt.Printf("Clone Object 's type: %T\n", cloneObj)
}
