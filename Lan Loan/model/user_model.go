package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Nim      string `json:"nim" form:"nim"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Class    string `json:"class" form:"class"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}
