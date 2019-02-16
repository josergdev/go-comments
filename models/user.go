package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Username        string    `json:"username" gorm:"not null;unique"`
	Email           string    `json:"email" form:"not null;unique"`
	FullName        string    `json:"fullName" gorm:"not null"`
	Password        string    `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	ConfirmPassword string    `json:"confirmPassword,omitempty" gorm:"-"`
	Picture         string    `json:"picture"`
	Comments        []Comment `json:"comments,omitempty"`
}
