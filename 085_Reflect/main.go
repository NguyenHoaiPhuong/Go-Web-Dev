package main

// https://blog.golang.org/laws-of-reflection

import (
	"fmt"
	"reflect"
)

// TestT1 : reflect.TypeOf and reflect.ValueOf, reflect.KindOf
func TestT1() {
	x := 6.4
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	k := v.Kind()
	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
	fmt.Println("Value:", v.String())
	fmt.Println("Kind:", k)
	// fmt.Printf("%T\n", x)
}

// TestT2 : reflect.Value 's methods Type(), Kind()
func TestT2() {
	x := 3.4
	v := reflect.ValueOf(x)
	fmt.Println("Type:", v.Type())
	fmt.Println("Kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("Value:", v.Float())
	// fmt.Println("Value:", v.Int())
}

// TestT3 : Settability & modify reflect.Value
func TestT3() {
	x := 3.4
	fmt.Println("x =", x)
	v := reflect.ValueOf(x)
	fmt.Println("v is Settable?", v.CanSet())

	p := reflect.ValueOf(&x)
	fmt.Println("Type of p:", p.Type())
	fmt.Println("p is Settable?", p.CanSet())
	fmt.Println("*p is Settable?", p.Elem().CanSet())
	p.Elem().SetFloat(7.7)
	fmt.Println("x =", x)
}

// TestT4 : settabilities of a struct 's fields value
func TestT4() {
	type human struct {
		age  int
		name string
	}
	akagi := human{age: 34, name: "Akagi"}
	v := reflect.ValueOf(akagi)
	fmt.Println("Value of akagi:", v)
	fmt.Println("Type of akagi:", v.Type())
	fmt.Println("Kind of akagi:", v.Kind())
	fmt.Println("akagi is Settable?:", v.CanSet())
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Println("Field", i)
		fmt.Println("Value of current field:", f)
		fmt.Println("Type of current field:", f.Type())
		fmt.Println("Kind of current field:", f.Kind())
		fmt.Println("Current field is Settable?:", f.CanSet())
	}
}

// TestT5 : modify value of struct 's fields
func TestT5() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}

func main() {
	// TestT1()
	// TestT2()
	// TestT3()
	// TestT4()
	TestT5()
}
