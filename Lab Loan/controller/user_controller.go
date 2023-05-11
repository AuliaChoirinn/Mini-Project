package controller

import (
	"lab-loan/config"
	"lab-loan/middleware"
	"lab-loan/model"
	"net/http"
	"github.com/labstack/echo/v4"
)

// User_register
func User_register(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to register user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

// User_login
func User_login(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	c.Get("user")

	// err_cek1
	err_cek1 := config.DB.Where("email=? AND password=?", user.Email, user.Password).First(&user).Error
	if err_cek1 != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "wrong email or password",
			"error":   err_cek1.Error(),
		})
	}

	// err_cek2
	token, err_cek2 := middleware.CreateTokenUser(user.Email, user.Password, user.Role)
	if err_cek2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to generate tokens",
			"error":   err_cek2.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "successfully logged in",
		"username": user.Username,
		"email":    user.Email,
		"token":    token,
		"nim":      user.Nim,
		"class":    user.Class,
		"role":     user.Role,
		"status":   user.Status,
	})
}
