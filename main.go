package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
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
	newInventoryRepository(db)
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

type inventoryRepository struct {
	db *gorm.DB
}

func newInventoryRepository(db *gorm.DB) {
	repo := new(inventoryRepository)
	repo.db = db
	repo.run()
}
func (i *inventoryRepository) run() {
	//category, err := i.CreateCategory(Category{
	//	CategoryName: "Sembako",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(category.ToString())

	//product, err := i.CreateProduct(Product{
	//	ProductCode: "A0002",
	//	ProductName: "Minyak goreng",
	//	CategoryId:  "0e766440-6ab6-43fa-b9c2-e5211039baa8",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(product.ToString())

	//Create product plus create category
	//product, err := i.CreateProduct(Product{
	//	ProductCode: "A0003",
	//	ProductName: "Sabun Batang",
	//	Category: Category{
	//		CategoryName: "Peralatan Mandi",
	//	},
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(product.ToString())
	//fmt.Printf("%s\n", strings.Repeat("=", 50))
	//products, err := i.getProductWithCategory()
	//if err != nil {
	//	panic(err)
	//}
	//for _, product := range products {
	//	fmt.Println(product.ToString())
	//}

	fmt.Printf("%s\n", strings.Repeat("=", 50))
	categories, err := i.getCategories()
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.ToString())
	}
}
func (i *inventoryRepository) getCategories() ([]Category, error) {
	categories := make([]Category, 0)
	err := i.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (i *inventoryRepository) getProductWithCategory() ([]Product, error) {
	products := make([]Product, 0)
	err := i.db.Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (i *inventoryRepository) CreateCategory(category Category) (*Category, error) {
	if err := i.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (i *inventoryRepository) CreateProduct(product Product) (*Product, error) {
	if err := i.db.Create(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
