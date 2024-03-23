package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name	string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaUrl	string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~SocialMediaUrl is required"`
	UserID uint
	User *User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}