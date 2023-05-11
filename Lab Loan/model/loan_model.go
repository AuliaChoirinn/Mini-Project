package model

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	Id_user     int       `json:"id_user" form:"id_user"`
	Id_tool     int       `json:"id_tool" form:"id_tool"`
	Loan_date   time.Time `json:"loan_date" form:"loan_date"`
	Return_date time.Time `json:"return_date" form:"return_date"`
}
