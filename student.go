package main

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID       int `gorm:"type:int4;autoIncrement"`
	Name     string
	Gender   string
	Age      int
	JoinDate time.Time
	IdCard   string
	Senior   bool
	gorm.Model
}

func (p *Student) TableName() string {
	return "m_student"
}
func (p *Student) ToString() string {
	student, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(student)
}
