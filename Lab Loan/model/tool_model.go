package model

import "gorm.io/gorm"


type Tool struct {
	gorm.Model
	Name          string `json:"name" form:"name"`
	Serial_number string `json:"serial_number" form:"serial_number"`
	Stock         int    `json:"stock" form:"stock"`
}
