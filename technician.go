package main

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Technician struct {
	ID                  uint
	TechnicianFirstName string
	TechnicianLastName  string
	Age                 uint
	Sites               []*Site `gorm:"many2many:technician_sites;"`
	gorm.Model
}

func (p *Technician) ToString() string {
	technician, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(technician)
}
