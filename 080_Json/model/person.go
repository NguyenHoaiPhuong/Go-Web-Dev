package model

// Person includes name, sex, age
type Person struct {
	Name *PersonName `json:"name"`
	Sex  string      `json:"sex"`
	Age  int         `json:"age"`
}

// PersonName includes first, last and middle names
type PersonName struct {
	First  string `json:"first"`
	Last   string `json:"last"`
	Middle string `json:"middle"`
}
