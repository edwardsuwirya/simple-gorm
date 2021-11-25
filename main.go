package main

import (
	"database/sql"
	"errors"
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
	err = db.AutoMigrate(&Student{}, &Product{}, &Category{}, &UserInfo{}, &UserCredential{})
	enigmaDb, err := db.DB()
	defer func(enigmaDb *sql.DB) {
		err := enigmaDb.Close()
		if err != nil {
			panic(err)
		}
	}(enigmaDb)
	//err = enigmaDb.Ping()
	//if err != nil {
	//	panic(err)
	//}
	err = db.AutoMigrate(&Student{}, &Product{}, &Category{})
	if err != nil {
		panic(err)
	}
	//newStudentRepository(db)
	//newInventoryRepository(db)
	newUserInfoRepository(db)
}

type userInfoRepository struct {
	db *gorm.DB
}

func newUserInfoRepository(db *gorm.DB) {
	repo := new(userInfoRepository)
	repo.db = db
	repo.run()
}
func (r *userInfoRepository) run() {
	//newUser, err := r.CreateUser(UserInfo{
	//	FirstName: "Tika",
	//	LastName:  "Yesi",
	//	IdCard:    "922-933",
	//	Gender:    "F",
	//	Email:     "tika.yesi@corp.com",
	//	UserCredential: UserCredential{
	//		UserName:     "tika.yesi",
	//		UserPassword: "889911",
	//		IsBlocked:    false,
	//	},
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(newUser.ToString())
	//isUserExist, _ := r.UserValidation("berty.tanasale", "12345")
	//
	//fmt.Println("is valid", isUserExist)

	//userInfo, err := r.FindUserById(3)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		fmt.Println("Not found")
	//		return
	//	}
	//	panic(err)
	//}
	//fmt.Println(userInfo.ToString())
	//userInfo, err := r.FindUserById(3)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		fmt.Println("Not found")
	//		return
	//	}
	//	panic(err)
	//}
	//fmt.Println(userInfo.ToString())
	//userInfo, err := r.FindUserByCondition(UserInfo{
	//	Gender: "M",
	//})
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		fmt.Println("Not found")
	//		return
	//	}
	//	panic(err)
	//}
	//userInfo, err := r.FindUserDocument(4)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		fmt.Println("Not found")
	//		return
	//	}
	//	panic(err)
	//}
	//fmt.Println(userInfo)
	totalUserByGender, err := r.TotalUserGroupByGender()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Not found")
			return
		}
		panic(err)
	}
	fmt.Println(totalUserByGender)
}
func (r *userInfoRepository) CreateUser(user UserInfo) (*UserInfo, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userInfoRepository) FindUserById(id int) (*UserInfo, error) {
	var user UserInfo
	if err := r.db.First(&user, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type userDoc struct {
	IdCard string
	Email  string
}

func (r *userInfoRepository) FindUserByCondition(useCriteria UserInfo) ([]userDoc, error) {
	var userDoc []userDoc
	err := r.db.Model(&UserInfo{}).Where(&useCriteria).Find(&userDoc).Error
	if err != nil {
		return nil, err
	}
	return userDoc, nil
}

func (r *userInfoRepository) FindUserDocument(id int) ([]userDoc, error) {
	var userDoc []userDoc
	err := r.db.Table("user_infos").Select("id_card", "email").Where("id=?", id).Scan(&userDoc).Error
	if err != nil {
		return nil, err
	}
	return userDoc, nil
}

type Result struct {
	Gender string
	Total  int64
}

func (r *userInfoRepository) TotalUserGroupByGender() ([]Result, error) {
	var result []Result
	err := r.db.Table("user_infos").Select("gender", "count(id) as total").Group("gender").Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userInfoRepository) UserValidation(userName string, password string) (bool, error) {
	var userCount int64
	err := r.db.Table("user_credentials").Where("user_name=? AND user_password=? AND is_blocked=false", userName, password).Count(&userCount).Error
	if err != nil {
		return false, err
	}
	if userCount == 1 {
		return true, nil
	}
	return false, nil
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
	fmt.Printf("%s\n", strings.Repeat("=", 50))
	products, err := i.getProductWithCategory()
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Println(product.ToString())
	}

	//fmt.Printf("%s\n", strings.Repeat("=", 50))
	//categories, err := i.getCategories()
	//if err != nil {
	//	panic(err)
	//}
	//for _, category := range categories {
	//	fmt.Println(category.ToString())
	//}
}

func (i *inventoryRepository) getCategories() ([]Category, error) {
	categories := make([]Category, 0)
	//err := i.db.Find(&categories).Error
	err := i.db.Preload("Products").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

type pp struct {
	ProductCode string
}

func (i *inventoryRepository) getProductWithCategory() ([]Product, error) {
	//category := make([]Category, 0)
	products := make([]Product, 0)

	//var category Category
	//err := i.db.Debug().Joins("Category", i.db.Where(&Category{CategoryName: "Sembako"})).Find(&products).Error
	err := i.db.Debug().Joins("JOIN m_category on m_category.id=m_product.category_id AND m_category.category_name=?", "Sembako").Find(&products).Error
	//err := i.db.Debug().Table("m_product").Select("product_code").Where("category_id=?", "0e766440-6ab6-43fa-b9c2-e5211039baa8").Scan(&products).Error
	//category:=[]Category{
	//	{
	//		ID:           "0e766440-6ab6-43fa-b9c2-e5211039baa8",
	//	},
	//}
	//err := i.db.Debug().Model(category).Association("Products").Find(&products)
	//err := i.db.Debug().Preload("Products","product_code=?","A0002").Find(&category).Error
	//err:=i.db.Debug().Table("m_category").Select("category_name").Scan(&category).Error
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
