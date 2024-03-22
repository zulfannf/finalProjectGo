package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title	string `gorm:"not null" json:"title" form:"title" valid:"required~Title of your product is required"`
	Caption	string 
	PhotoUrl string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo_url of your product is required"`
	UserId uint
	User *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}