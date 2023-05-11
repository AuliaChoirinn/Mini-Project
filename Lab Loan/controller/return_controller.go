package controller

import (
	"lab-loan/config"
	"lab-loan/model"
	"net/http"
	"time"
	"github.com/labstack/echo/v4"
)

// Add_return
func Add_returning(c echo.Context) error {
	returning := model.Return{}
	c.Bind(&returning)

	// cek_err1 : apakah id_loan exist
	var loan model.Loan
	err_cek1 := config.DB.First(&loan, returning.Id_loan).Error
	if err_cek1 != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id_loan not found",
			"error":   err_cek1.Error(),
		})
	}

	if loan.Id_user!=returning.Id_user{
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id_user not found in current loan record",
			"error":   err_cek1.Error(),
		})
	}

	// cek_err2 : apakah id_tool exist di row tabel loan
	var tool model.Tool
	err_cek2 := config.DB.First(&tool, returning.Id_tool).Error
	if err_cek2 != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id_tool not found",
			"error":   err_cek2.Error(),
		})
	}

	today := time.Now()
	returning.Return_date = today

	var status string
	if loan.Loan_date.After(today) {
		status = "late"
	} else {
		status = "ontime"
	}
	returning.Status = status

	err_cek3 := config.DB.Save(&returning).Error
	if err_cek3 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to add returning",
			"error":   err_cek3.Error(),
		})
	}

	// menambah stock karena barang telah dikembalikan
	tool.Stock += 1

	err_cek4 := config.DB.Model(&model.Tool{}).Where("id=?", returning.Id_tool).Updates(map[string]interface{}{"stock": tool.Stock}).Error
	if err_cek4 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to increment stock",
			"error":   err_cek4.Error(),
		})
	}

	config.DB.Delete(&loan)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully added the returning",
		"return":  returning,
	})
}

// Get_returnings
func Get_returnings(c echo.Context) error {
	var user model.User
	c.Bind(&user)
	
	var returning []model.Return
	err := config.DB.Where("id_user=?", user.ID).Find(&returning).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get returns",
		"return":  returning,
	})
}
