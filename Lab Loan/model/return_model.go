package model

import (
	"time"

	"gorm.io/gorm"
)

type Return struct {
	gorm.Model
	Id_loan     int       `json:"id_loan" form:"id_loan"`
	Id_user     int       `json:"id_user" form:"id_user"`
	Id_tool     int       `json:"id_tool" form:"id_tool"`
	Return_date time.Time `json:"return_date" form:"return_date"`
	Status      string    `json:"status" form:"status"`
}
