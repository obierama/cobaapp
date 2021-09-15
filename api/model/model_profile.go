package model

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	Alamat_ktp          string `gorm:"size:100;not null" json:"alamat_ktp"`
	Pekerjaan           string `gorm:"size:100;" json:"pekerjaan"`
	Nama_lengkap        string `gorm:"size:100;not null;" json:"nama_lengkap"`
	Pendidikan_terakhir string `gorm:"size:100;not null;" json:"pendidikan_terakhir"`
	Nomer_tlp           string `gorm:"size:100;not null;" json:"nomer_tlp"`
	//untuk join tabel
	Userid int64 `gorm:"" json:"userid"`
}
type StukturProfile struct {
	Id                  int64  `json:"id"`
	Alamat_ktp          string `json:"alamat_ktp"`
	Pekerjaan           string `json:"pekerjaan"`
	Nama_lengkap        string `json:"nama_lengkap"`
	Pendidikan_terakhir string `json:"pendidikan_terakhir"`
	Nomer_tlp           string `json:"nomer_tlp"`
	//untuk join tabel
	Name  string `json:"name"`
	Email string `json:"email"`
}

//untuk join tabel  32-44
var (
	seprofile = "profiles.*, users.name , users.email "
	joinuser  = "JOIN users on profiles.userid = users.id"
)

func (u *Profile) FindInt(db *gorm.DB, param string) (*[]StukturProfile, error) {
	data := []StukturProfile{}
	if err := db.Debug().Table("profiles").Select(seprofile).Joins(joinuser).Where("users.name ILike ?", "%"+param+"%").Find(&data).Error; err != nil { //tambah "%"+nama_param+"%" hanya untuk menambah fungsi searching
		return &[]StukturProfile{}, err
	}
	return &data, nil
}

func (v *Profile) Update(id uint, db *gorm.DB) (*Profile, error) {
	//hashedpassword, _ := HashPassword(v.Password)
	if err := db.Debug().Table("profiles").Where("id = ?", id).Updates(Profile{
		Alamat_ktp:          v.Alamat_ktp,
		Pekerjaan:           v.Pekerjaan,
		Nama_lengkap:        v.Nama_lengkap,
		Pendidikan_terakhir: v.Pendidikan_terakhir,
		Nomer_tlp:           v.Nomer_tlp,
	}).Error; err != nil {
		return &Profile{}, err
	}
	return v, nil
}

func (u *Profile) Validate(action string) error {
	switch strings.ToLower(action) {

	case "update":

		if u.ID == 0 {
			return errors.New("id is required")
		}
		if u.Alamat_ktp == "" {
			return errors.New("alamat_ktp is required")
		}
		if u.Pekerjaan == "" {
			return errors.New("pekerjaan is required")
		}

		if u.Nama_lengkap == "" {
			return errors.New("nama_lengkap is required")
		}

		if u.Pendidikan_terakhir == "" {
			return errors.New("pendidikan_terakhir is required")
		}

		if u.Nomer_tlp == "" {
			return errors.New("nomer_tlp is required")
		}

		return nil

	default:
		if u.Alamat_ktp == "" {
			return errors.New("alamat_ktp is required")
		}
		if u.Pekerjaan == "" {
			return errors.New("pekerjaan is required")
		}

		if u.Nama_lengkap == "" {
			return errors.New("nama_lengkap is required")
		}

		if u.Pendidikan_terakhir == "" {
			return errors.New("pendidikan_terakhir is required")
		}

		if u.Nomer_tlp == "" {
			return errors.New("nomer_tlp is required")
		}

		//untuk join tabel
		if u.Userid == 0 {
			return errors.New("user id is required")
		}

		return nil

	}

}

func (u *Profile) SaveProfile(db *gorm.DB) (*Profile, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Profile{}, err
	}
	return u, nil
}

func (u *Profile) GetProfile(db *gorm.DB, parameter string, data string) (*Profile, error) {
	account := &Profile{}
	if err := db.Debug().Table("profiles").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (u *Profile) GetProfileInt(db *gorm.DB, parameter string, data int) (*Profile, error) {
	account := &Profile{}
	if err := db.Debug().Table("profiles").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

//untuk fungsu delete/hapus
func (u *Profile) Delete(db *gorm.DB, data string) (*Profile, error) {
	account := &Profile{}
	if err := db.Unscoped().Table("profiles").Where("id = ?", data).Delete(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
