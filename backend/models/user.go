package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nombre   string `json:"nombre"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
