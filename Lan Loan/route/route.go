package route

import (
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"lab-loan/constant"
	"lab-loan/controller"
)

func New() *echo.Echo {
	e := echo.New()

	// for admin
	e.GET("/check", controller.Check)
	e.POST("/admin/login", controller.Admin_login)

	// for user
	e.POST("/user/register", controller.User_register)
	e.POST("/user/login", controller.User_login)

	// for user and admin
	e.GET("/tools", controller.Get_tools)

	// authorization token jwt
	// for admin
	authJWTAdmin := e.Group("")
	authJWTAdmin.Use(middleware.JWT([]byte(constant.SECRET_JWT_ADMIN)))
	authJWTAdmin.POST("/admin/tool", controller.Add_tool)
	authJWTAdmin.DELETE("/admin/tool/:id", controller.Delete_tool)

	// for user
	authJWTUser := e.Group("")
	authJWTUser.Use(middleware.JWT([]byte(constant.SECRET_JWT_USER)))
	authJWTUser.POST("/user/loan", controller.Add_loan)
	authJWTUser.GET("/user/loan", controller.Get_loans)
	authJWTUser.POST("/user/return", controller.Add_returning)
	authJWTUser.GET("/user/return", controller.Get_returnings)

	return e
}
