package model

// Person includes name, sex
type Person struct {
	Name *PersonName `json:"name"`
	Sex  string      `json:"sex"`
}

// Human includes name, sex
type Human struct {
	Name *HumanName
	Sex  string
}

// PersonName includes first, last and middle names
type PersonName struct {
	First  string `json:"first"`
	Last   string `json:"last"`
	Middle string `json:"middle"`
}

// HumanName includes first, last and middle names
type HumanName struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Mid   string `json:"middle"`
}

// Person includes name, sex, age
type PersonFull struct {
	Name *PersonName `json:"name"`
	Sex  string      `json:"sex"`
	Age  int         `json:"age"`
}
