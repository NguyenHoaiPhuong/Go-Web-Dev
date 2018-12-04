package test

import (
	"GO-WEB-DEV/082_Go-Model-By-Jeevatkm/product"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"gopkg.in/jeevatkm/go-model.v1"
)

func TestAddConversion(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestAddConversion: String to Int")

	model.AddConversion((*string)(nil), (*int)(nil), func(in reflect.Value) (reflect.Value, error) {
		intVal, err := strconv.Atoi(in.String())
		if err != nil {
			t.Errorf("Add Conversion string to int Error: %v", err)
			return reflect.Value{}, err
		}
		return reflect.ValueOf(intVal), nil
	})

	A := &product.A{
		ID:   "1234",
		Name: "Product A",
	}
	F := new(product.F)
	errs := model.Copy(F, A)
	if len(errs) > 0 {
		t.Errorf("%v\n", errs)
	}
	fmt.Printf("Source: %v\n", *A)
	fmt.Printf("Destination: %v\n", *F)

	model.RemoveConversion((*string)(nil), (*int)(nil))
}
