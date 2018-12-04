package product

// A product
type A struct {
	ID   string
	Name string
}

// B product
type B struct {
	ID          string
	Name        string
	ExpiredDate string
}

// C product
type C struct {
	ID   string
	Name string
	Code string
}

// D product
type D struct {
	Name []string
}

// E product
type E struct {
	Name   string `json:"name,omitempty"`
	Region string `gorm:"column:region;unique" json:"region,omitempty"`
}

// F product
type F struct {
	ID   int
	Name string
}
