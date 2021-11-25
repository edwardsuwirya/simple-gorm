package main

type UserCredential struct {
	ID           int
	UserName     string `gorm:"index:idx_user_name,unique"`
	UserPassword string
	IsBlocked    bool
}
