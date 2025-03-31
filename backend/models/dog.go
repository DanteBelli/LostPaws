package models

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	UserId      uint   `json:"user_id"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Location    string `json:"location"`
	Status      string `json:"status"`
}
