package controller

import (
	"github.com/labstack/echo/v4"
	"lab-loan/config"
	"lab-loan/middleware"
	"lab-loan/model"
	"net/http"
)

// Check
func Check(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Welcome to Lab Loan",
		"owner":   "Firman",
	})
}

// Admin_login
func Admin_login(c echo.Context) error {
	admin := model.Admin{}
	c.Bind(&admin)
	c.Get("user")

	// cek1
	err_cek1 := config.DB.Where("username=? AND password=?", admin.Username, admin.Password).First(&admin).Error
	if err_cek1 != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "wrong username or password",
			"error":   err_cek1.Error(),
		})
	}

	token, err_cek2 := middleware.CreateTokenAdmin(admin.Username, admin.Password, admin.Role)
	if err_cek2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "there was an error when generating the token",
			"error":   err_cek2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "successfully logged in",
		"username": admin.Username,
		"role":     admin.Role,
		"token":    token,
	})
}
