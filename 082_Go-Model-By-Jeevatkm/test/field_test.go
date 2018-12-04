package test

import (
	"GO-WEB-DEV/082_Go-Model-By-Jeevatkm/product"
	"fmt"
	"testing"

	model "gopkg.in/jeevatkm/go-model.v1"
)

func TestIsZero(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestIsZero")
	A := &product.A{
		ID:   "1234",
		Name: "",
	}
	B := &product.B{}
	fmt.Printf("Is A a zero struct? %v\n", model.IsZero(A))
	fmt.Printf("Is B a zero struct? %v\n", model.IsZero(B))
}

////////////////////////////////////////////////////////
func TestHasZero(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestHasZero")
	A := &product.A{
		ID:   "1234",
		Name: "",
	}
	fmt.Printf("A has field without value? %v\n", model.HasZero(A))
}

////////////////////////////////////////////////////////
func TestIsZeroInFields(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestIsZeroInFields")
	C := &product.C{
		ID:   "",
		Name: "",
		Code: "4375ABXD",
	}
	fieldName, isEmpty := model.IsZeroInFields(C, "ID", "Name", "Code") // return only the first field with zero value
	fmt.Println("Empty Field Name:", fieldName)
	fmt.Println("Yes, I have zero value:", isEmpty)

	fieldName, isEmpty = model.IsZeroInFields(C, "ID")
	fmt.Printf("%v Field is Empty? %v\n", fieldName, isEmpty)
	fieldName, isEmpty = model.IsZeroInFields(C, "Name")
	fmt.Printf("%v Field is Empty? %v\n", fieldName, isEmpty)
	fieldName, isEmpty = model.IsZeroInFields(C, "Code")
	fmt.Printf("%v Field is Empty? %v\n", fieldName, isEmpty)
}

////////////////////////////////////////////////////////
func TestFields(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestFields")
	C := &product.C{
		ID:   "",
		Name: "",
		Code: "4375ABXD",
	}
	fields, err := model.Fields(C)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	fmt.Println("Fields:", fields)
}

////////////////////////////////////////////////////////
func TestKind(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestKind")
	C := &product.C{
		ID:   "",
		Name: "",
		Code: "4375ABXD",
	}
	fieldKind, err := model.Kind(C, "Code")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	fmt.Printf("%v\n", fieldKind)
}

////////////////////////////////////////////////////////
func TestTag(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestTag")
	E := &product.E{}
	tag, err := model.Tag(E, "Name")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	fmt.Println("Tag Value:", tag.Get("json"))
	tag, err = model.Tag(E, "Region")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	fmt.Println("Tag Value:", tag.Get("gorm"))
}

////////////////////////////////////////////////////////
func TestTags(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestTags")
	E := &product.E{}
	tags, err := model.Tags(E)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	fmt.Println("Tags:", tags)

	tag := tags["Name"]
	fmt.Println("Tag Value:", tag.Get("json"))

	tag = tags["Region"]
	fmt.Println("Tag Value:", tag.Get("gorm"))

	tag = tags["Region"]
	fmt.Println("Tag Value:", tag.Get("json"))
}

////////////////////////////////////////////////////////
func TestGet(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestGet")
	C := &product.C{
		ID:   "1234",
		Name: "Hair Dryer",
		Code: "4375ABXD",
	}
	value, err := model.Get(C, "ID")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	fmt.Printf("Value: %v\n", value)
}

////////////////////////////////////////////////////////
func TestSet(t *testing.T) {
	fmt.Println("--------------------------------------")
	fmt.Println("TestSet")
	C := &product.C{
		ID:   "1234",
		Name: "Hair Dryer",
		Code: "4375ABXD",
	}
	value, _ := model.Get(C, "ID")
	fmt.Printf("Old Value: %v\n", value)
	err := model.Set(C, "ID", "4321")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	value, _ = model.Get(C, "ID")
	fmt.Printf("New Value: %v\n", value)
}
