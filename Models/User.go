package models

import (
	initializers "Basic/Auth-Api/Initializers"
	token "Basic/Auth-Api/Token"
	"encoding/json"
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FullName   string `json:"fulllName" binding:"required"`
	Password   string `json:"password" binding:"required,min=8,max=15"`
	Username   string `gorm:"unique" json:"userName" binding:"required,alpha"`
	ProfileImg string `json:"profileImg,omitempty" binding:"required"`
	RoleId     int    `json:"roleId" binding:"required"`
	Role       Role   `json:"role"`
	Status     int8   `json:"status" binding:"required"`
}

func GetUserByID(uid int) (User, error) {

	var u User

	if err := initializers.DB.Preload("Role").First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	return u, nil

}
func (u *User) PrepareGive() {
	u.Password = ""
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = initializers.DB.Model(User{}).Where("username = ?", username).Preload("Role").Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil { //err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID, u.Role.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = initializers.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func IsExist(username string) bool {
	var err error

	user := User{}

	err = initializers.DB.Model(User{}).Where("username = ?", username).Preload("Role").Take(&user).Error
	if err == nil && &user != nil {
		return true
	} else {
		return false
	}
}

func (u *User) GetUserById(id int) (*User, error) {
	var user User
	cashedUser, err := initializers.Client.Get(strconv.Itoa(id)).Result()
	if err == nil {
		json.Unmarshal([]byte(cashedUser), &user)
	} else {
		initializers.DB.Preload("Role").First(&user, id)
		strUser, err := json.Marshal(&user)
		if err == nil {
			initializers.Client.Set(string(strconv.Itoa(id)), &strUser, 30)
		}
	}
	return &user, nil
}

func (u *User) BeforeSave(db *gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil

}

type Status int

const (
	Active   Status = 1
	InActive Status = 2
)

func (s Status) String() string {
	return [...]string{"Active", "InActive"}[s]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (s Status) EnumIndex() int {
	return int(s)
}
