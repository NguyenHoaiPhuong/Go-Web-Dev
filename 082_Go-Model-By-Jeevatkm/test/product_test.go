package test

import (
	"GO-WEB-DEV/082_Go-Model-By-Jeevatkm/product"
	"fmt"
	"testing"

	model "gopkg.in/jeevatkm/go-model.v1"
)

func TestCopyProductAtoB(t *testing.T) {
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
