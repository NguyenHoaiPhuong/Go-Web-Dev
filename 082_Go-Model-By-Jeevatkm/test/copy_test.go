package test

import (
	"GO-WEB-DEV/082_Go-Model-By-Jeevatkm/product"
	"fmt"
	"testing"

	model "gopkg.in/jeevatkm/go-model.v1"
)

func TestCopyProductAtoB(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestCopyProductAtoB")
	A := &product.A{
		ID:   "1234",
		Name: "Hair Dry",
	}
	B := new(product.B)
	errs := model.Copy(B, A)
	if len(errs) > 0 {
		t.Errorf("%v\n", errs)
	}
	fmt.Printf("Source: %v\n", *A)
	fmt.Printf("Destination: %v\n", *B)
}

func TestCopyProductBtoA(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestCopyProductBtoA")
	B := &product.B{
		ID:          "1234",
		Name:        "Hair Dry",
		ExpiredDate: "2019-01-01",
	}
	A := new(product.A)
	errs := model.Copy(A, B)
	if len(errs) > 0 {
		t.Errorf("%v\n", errs)
	}
	fmt.Printf("Source: %v\n", *B)
	fmt.Printf("Destination: %v\n", *A)
}

func TestCopyProductBtoC(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestCopyProductBtoC")
	B := &product.B{
		ID:          "1234",
		Name:        "Hair Dry",
		ExpiredDate: "2019-01-01",
	}
	C := new(product.C)
	errs := model.Copy(C, B)
	if len(errs) > 0 {
		t.Errorf("%v\n", errs)
	}
	fmt.Printf("Source: %v\n", *B)
	fmt.Printf("Destination: %v\n", *C)
}

func TestCopyProductCtoD(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestCopyProductCtoD")
	C := &product.C{
		ID:   "1234",
		Name: "Hair Dry",
		Code: "ABC123XUV",
	}
	D := new(product.D)
	errs := model.Copy(D, C)
	if len(errs) > 0 {
		t.Errorf("%v\n", errs)
	}
	fmt.Printf("Source: %v\n", *C)
	fmt.Printf("Destination: %v\n", *D)
}

func TestCopyProductDtoC(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestCopyProductDtoC")
	D := &product.D{
		Name: []string{"Hair Dry", "Machine"},
	}
	C := new(product.C)
	errs := model.Copy(C, D)
	if len(errs) > 0 {
		t.Errorf("%v\n", errs)
	}
	fmt.Printf("Source: %v\n", *D)
	fmt.Printf("Destination: %v\n", *C)
}
