package main

import (
	"encoding/json"
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string   `gorm:"column:id;size:36;primaryKey"`
	ProductCode string   `gorm:"column:product_code;size:255;not null;index:uq_product_code,unique"`
	ProductName string   `gorm:"column:product_name;size:255;not null"`
	CategoryId  string   `gorm:"column:category_id;size:36;not null"`
	Category    Category `gorm:"foreignkey:CategoryId"`
	gorm.Model
}

func (p *Product) TableName() string {
	return "m_product"
}
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = guuid.New().String()
	return nil
}
func (p *Product) ToString() string {
	product, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(product)
}
