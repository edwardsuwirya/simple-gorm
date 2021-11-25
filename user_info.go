package main

import (
	"encoding/json"
	"gorm.io/gorm"
)

type UserInfo struct {
	FirstName      string
	LastName       string
	IdCard         string
	Gender         string
	Email          string `gorm:"index:idx_user_email,unique"`
	CredentialID   int
	UserCredential UserCredential `gorm:"foreignKey:CredentialID"`
	gorm.Model
}

func (p *UserInfo) ToString() string {
	student, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(student)
}
