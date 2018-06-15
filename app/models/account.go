package models

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
)

type Priviledge int

type User struct {
	ID         int        `json:"-"`
	Name       string     `sql:"unique;not null" json:"name"`
	Password   []byte     `sql:"not null" json:"-"`
	Email      string     `sql:"unique;not null" json:"email,omitempty"`
	Score      int        `sql:"-" json:"score,omitempty"`
	AuthLevel  int        `sql:"not null" json:"-"`
	IsITF      bool       `json:"is_itf,omitempty"`
	Comment    string     `json:"comment"`
	CreatedAt  *time.Time `json:"-"`
	DeletedAt  *time.Time `sql:"index" json:"-"`
}

/*
AuthLevel
100  : Normal User
0 : Admin User
*/

var AuthLevelGroup = struct {
	Admin  int
	Normal int
}{
	0,
	100,
}

func GenPasswordHash(password string) []byte {

	//casting string to byte causes memory copy
	//so if it is buttle neck, consider using unsafe

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	//hash, err := bcrypt.GenerateFromPassword(*(*[]byte)(unsafe.Pointer(&password)), 10)

	if err != nil {
		return nil
	}

	return hash
}

func NewAdminUser(name string, email string, password string, isITF bool) *User {

	hash := GenPasswordHash(password)

	if hash == nil {
		return &User{}
	}

	return &User{
		Name:       name,
		Email:      email,
		Password:   hash,
		Score:      0,
		IsITF:      isITF,
		AuthLevel:  AuthLevelGroup.Admin,
		DeletedAt:  nil,
	}
}

func NewUser(name string, email string, password string, isITF bool) *User {

	hash := GenPasswordHash(password)

	if hash == nil {
		return &User{}
	}

	return &User{
		Name:       name,
		Email:      email,
		Password:   hash,
		Score:      0,
		IsITF:      isITF,
		AuthLevel:  AuthLevelGroup.Normal,
		DeletedAt:  nil,
	}
}

func IsAdmin(dbsess *gorm.DB, username string) bool {

	var user User

	err := dbsess.Table("users").Where("name = ?", username).First(&user).Error
	if err != nil {
		log.Println(err)
		return false
	}

	if user.AuthLevel != AuthLevelGroup.Admin {
		return false
	}

	return true

}

func CheckIfUserExistsByName(dbsess *gorm.DB, name string) bool {

	var user User

	dbsess.Table("users").Where("name = ?", name).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}

func RegisterUser(dbsess *gorm.DB, user *User) bool {

	if dbsess.Table("users").Create(user).Error != nil {
		return false
	}

	return true
}

func AuthUserByCred(dbsess *gorm.DB, name string, password string) int {

	var user User
	dbsess.Table("users").Where("name = ?", name).First(&user)

	if user.ID == 0 {
		return -1
	}

	if bcrypt.CompareHashAndPassword(user.Password, []byte(password)) != nil {
		return -1
	}

	return user.ID
}

func AddUser(dbsess *gorm.DB, user *User) {

	dbsess.Create(&user)
	dbsess.Save(&user)
}

func UpdateUserProfile(dbsess *gorm.DB, user *User) bool {

	err := dbsess.Table("users").Save(&user).Error
	if err != nil {
		log.Println(err)
		return false
	}

	return true	
	
}

func DeleteUserByID(dbsess *gorm.DB, userID int) bool {

	if dbsess.Table("users").Where("users.id = ?", userID).Delete(User{}).Error != nil {
		return false
	}
	return true
}

func GetUserIDByName(dbsess *gorm.DB, name string) int {

	var user User

	if dbsess.Table("users").Select("users.id").Where("users.name = ?", name).First(&user).Error != nil {
		return -1
	}

	return user.ID
}


//Get User Info inclding Password Hash
func GetUserByName(dbsess *gorm.DB, name string) User {
	
	var user User
	
	err := dbsess.Table("users").Where("users.name = ?", name).First(&user).Error
	if err != nil{
		log.Print(err)
		user.ID = -1
		return user
	}

	return user
}
	

func GetUserProfileByName(dbsess *gorm.DB, name string) User {

	var user User

	err := dbsess.Table("users").Select([]string{"users.id", "users.name", "users.email", "users.is_itf", "users.comment"}).Where("users.name = ?", name).First(&user).Error

	if err != nil{
		log.Println(err)
		user.ID = -1
		return user
	}

	
	return user
}


func GetUserProfileByID(dbsess *gorm.DB, userID int) User {

	var user User

	err := dbsess.Table("users").Select([]string{"users.id", "users.name", "users.is_itf", "users.comment"}).Where("users.id = ?", userID).First(&user).Error

	if err != nil{
		log.Println(err)
		user.ID = -1
		return user
	}

	
	return user
}



func GetOwnProfileByID(dbsess *gorm.DB, userID int) User {

	var user User

	err := dbsess.Table("users").Select([]string{"users.id", "users.name", "users.email", "users.is_itf", "users.comment"}).Where("users.id = ?", userID).First(&user).Error

	if err != nil{
		log.Println(err)
		user.ID = -1
		return user
	}

	
	return user
}
