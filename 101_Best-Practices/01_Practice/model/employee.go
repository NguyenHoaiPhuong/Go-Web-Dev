package model

import (
	"github.com/jinzhu/gorm"
)

// Employee definition
type Employee struct {
	gorm.Model
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
