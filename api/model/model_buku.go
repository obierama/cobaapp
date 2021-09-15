package model

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type Buku struct {
	gorm.Model
	Nama       string `gorm:"size:100;not null" json:"nama"`
	Judul_buku string `gorm:"size:100;" json:"judul_buku"`
	Penulis    string `gorm:"size:100;not null;unique_index" json:"penulis"`
}
type StukturBuku struct {
	Id         int64  `json:"id"`
	Nama       string `json:"nama"`
	Judul_buku string `json:"judul_buku"`
	Penulis    string `json:"penulis"`
}

func (u *Buku) Validate(action string) error {
	switch strings.ToLower(action) {

	case "update_buku":
		if u.Nama == "" {
			return errors.New("nama is required")
		}
		if u.Judul_buku == "" {
			return errors.New("judul_buku is required")
		}

		if u.Penulis == "" {
			return errors.New("penulis is required")
		}

		return nil

	default:
		if u.Nama == "" {
			return errors.New("nama is required")
		}
		if u.Judul_buku == "" {
			return errors.New("judul_buku is required")
		}

		if u.Penulis == "" {
			return errors.New("penulis is required")
		}

		return nil

	}

}

func (u *Buku) SaveBuku(db *gorm.DB) (*Buku, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Buku{}, err
	}
	return u, nil
}

func (u *Buku) GetBuku(db *gorm.DB, parameter string, data string) (*Buku, error) {
	account := &Buku{}
	if err := db.Debug().Table("bukus").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (u *Buku) GetBukuInt(db *gorm.DB, parameter string, data int) (*Buku, error) {
	account := &Buku{}
	if err := db.Debug().Table("bukus").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
