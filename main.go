package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func main() {
	dbHost := "159.223.42.164"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "P@ssw0rd"
	dbName := "enigma"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Student{}, &Product{}, &Category{})
	if err != nil {
		panic(err)
	}

	students, err := getAllStudent(db)
	if err != nil {
		panic(err)
	}
	for _, student := range students {
		fmt.Println(student.ToString())
	}
	fmt.Println("=========================================")
	students, err = GetStudentByName(db, "R")
	if err != nil {
		panic(err)
	}
	for _, student := range students {
		fmt.Println(student.ToString())
	}

	fmt.Println("=========================================")
	student := Student{
		Name:     "Pito",
		Gender:   "M",
		Age:      30,
		JoinDate: time.Time{},
		IdCard:   "333",
		Senior:   true,
	}
	newStudent, err := CreateStudent(db, student)
	if err != nil {
		panic(err)
	}
	fmt.Println(newStudent.ToString())

	fmt.Println("=========================================")
	err = DeleteStudent(db, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Delete")

}
func getAllStudent(db *gorm.DB) ([]Student, error) {
	students := make([]Student, 0)
	err := db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
func GetStudentByName(db *gorm.DB, name string) ([]Student, error) {
	students := make([]Student, 0)
	err := db.Where("Name LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
func CreateStudent(db *gorm.DB, student Student) (*Student, error) {
	err := db.Create(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func DeleteStudent(db *gorm.DB, id int) error {
	err := db.Delete(&Student{}, id).Error
	return err
}
