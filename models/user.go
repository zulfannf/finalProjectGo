package models

import (
	"mygram_finalprojectgo/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~username tidak boleh kosong"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~email tidak boleh kosong,email~email tidak tepat"`
	Age      int    `gorm:"not null;check:age > 8" json:"age" form:"age" valid:"required~umur harus diatas 8"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~ password tidak boleh kosong, minstringlength(6)~Password minimal 6 karakter"`
	Photos   []Photo
	Comments []Comment
	// SocialMedias [] SocialMedia
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

//validasi umur
func (umur *User) age() bool {
	return umur.Age > 8
}

