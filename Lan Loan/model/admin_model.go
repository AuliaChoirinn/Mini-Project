package model

type Admin struct {
	ID       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}
