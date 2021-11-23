package main

import (
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
