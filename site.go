package main

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Site struct {
	ID         uint
	SiteName   string
	Technician []*Technician `gorm:"many2many:technician_sites;"`
	gorm.Model
}

func (p *Site) ToString() string {
	site, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(site)
}
