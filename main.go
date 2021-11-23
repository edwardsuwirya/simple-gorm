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
	//newStudentRepository(db)

}

type studentRepository struct {
	db *gorm.DB
}

func newStudentRepository(db *gorm.DB) {
	repo := new(studentRepository)
	repo.db = db
	repo.run()
}
func (r *studentRepository) run() {
	students, err := r.getAllStudent()
	if err != nil {
		panic(err)
	}
	for _, student := range students {
		fmt.Println(student.ToString())
	}
	fmt.Println("=========================================")
	students, err = r.getStudentByName("R")
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
	newStudent, err := r.createStudent(student)
	if err != nil {
		panic(err)
	}
	fmt.Println(newStudent.ToString())

	fmt.Println("=========================================")
	err = r.deleteStudent(10)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Delete")
}
func (r *studentRepository) getAllStudent() ([]Student, error) {
	students := make([]Student, 0)
	err := r.db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
func (r *studentRepository) getStudentByName(name string) ([]Student, error) {
	students := make([]Student, 0)
	err := r.db.Where("Name LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
func (r *studentRepository) createStudent(student Student) (*Student, error) {
	err := r.db.Create(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
func (r *studentRepository) deleteStudent(id int) error {
	err := r.db.Delete(&Student{}, id).Error
	return err
}
