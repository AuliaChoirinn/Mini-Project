package controller

import (
	"lab-loan/config"
	"lab-loan/model"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

// Add_loan
func Add_loan(c echo.Context) error {
	loan := model.Loan{}
	c.Bind(&loan)

	// cek_err1 : apakah id_user exist
	var user model.User
	err_cek1 := config.DB.First(&user, loan.Id_user).Error
	if err_cek1 != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id_user not found",
			"error":   err_cek1.Error(),
		})
	}

	// cek_err1 : apakah id_tool exist
	var tool model.Tool
	err_cek2 := config.DB.First(&tool, loan.Id_tool).Error
	if err_cek2 != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id_tool not found",
			"error":   err_cek2.Error(),
		})
	}

	today := time.Now()
	loan.Loan_date = today
	loan.Return_date = today.AddDate(0, 0, 7)

	err_cek3 := config.DB.Save(&loan).Error
	if err_cek3 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to add loan",
			"error":   err_cek3.Error(),
		})
	}

	// mengurangi stock karena di pinjam
	tool.Stock -= 1

	err_cek4 := config.DB.Model(&model.Tool{}).Where("id=?", loan.Id_tool).Updates(map[string]interface{}{"stock": tool.Stock}).Error
	if err_cek4 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to decrement stock",
			"error":   err_cek4.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully added the loan",
		"loan":    loan,
	})
}

// Get_loans
func Get_loans(c echo.Context) error {
	var user model.User
	c.Bind(&user)
	var loan []model.Loan
	err := config.DB.Where("id_user=?", user.ID).Find(&loan).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get loans",
		"loan":    loan,
	})
}
