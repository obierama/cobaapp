package model

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	//"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name         string `gorm:"size:100;not null" json:"name"`
	Nama_lengkap string `gorm:"size:100;" json:"nama_lengkap"`
	Email        string `gorm:"size:100;not null;unique_index" json:"email"`
	Password     string `gorm:"size:100;not null" json:"password"`
}
type StukturUser struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Nama_lengkap string `json:"nama_lengkap"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}
	return nil
}

func (u *User) BeforeSave() error {
	password := strings.TrimSpace(u.Password)
	hashedpassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = string(hashedpassword)
	return nil
}

func (u *User) Prepare() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Email == "" {
			return errors.New("Email is required")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		return nil
	case "update":
		if u.Name == "" {
			return errors.New("Name is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}

		if u.Password == "" {
			return errors.New("Password is required")
		}

		if u.Nama_lengkap == "" {

			return errors.New("Nama Lengkap is required")
		}

		return nil

	default:
		if u.Name == "" {
			return errors.New("Name is required")
		}
		if u.Email == "" {
			return errors.New("Email is required")
		}

		if u.Password == "" {
			return errors.New("Password is required")
		}

		if u.Nama_lengkap == "" {

			return errors.New("Nama Lengkap is required")
		}

		return nil

	}

}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GetUser(db *gorm.DB, parameter string, data string) (*User, error) {
	account := &User{}
	if err := db.Debug().Table("users").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (u *User) GetUserInt(db *gorm.DB, parameter string, data int) (*User, error) {
	account := &User{}
	if err := db.Debug().Table("users").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
