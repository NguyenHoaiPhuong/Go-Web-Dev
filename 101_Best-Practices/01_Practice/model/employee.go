package model

// Employee definition
type Employee struct {
	ID     string `gorm:"primary_key" json:"id"`
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
}

// Disable sets status false
func (e *Employee) Disable() {
	e.Status = false
}

// Enable sets status true
func (e *Employee) Enable() {
	e.Status = true
}
