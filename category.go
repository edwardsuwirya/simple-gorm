package main

import (
	"encoding/json"
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID           string `gorm:"column:id;size:36;primaryKey"`
	CategoryName string `gorm:"column:category_name;size:255;not null;index:uq_category,unique"`
	gorm.Model
}

func (c *Category) TableName() string {
	return "m_category"
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = guuid.New().String()
	return nil
}
func (c *Category) ToString() string {
	category, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(category)
}
