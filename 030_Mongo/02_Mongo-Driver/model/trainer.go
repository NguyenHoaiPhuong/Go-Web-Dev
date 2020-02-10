package model

import "reflect"

// Trainer : a struct include Name, Age and City
type Trainer struct {
	Name string
	Age  int
	City string
}

type Trainers []*Trainer

func (trainers *Trainers) ConvertToInterfaceSlice() []interface{} {
	list := reflect.ValueOf(trainers).Elem()
	items := make([]interface{}, list.Len())
	for i := 0; i < list.Len(); i++ {
		items[i] = list.Index(i).Interface()
	}
	return items
}
