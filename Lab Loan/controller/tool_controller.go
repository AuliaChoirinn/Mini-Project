package controller

import (
	"lab-loan/config"
	"lab-loan/model"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

// Add_tool
func Add_tool(c echo.Context) error {

	// token := c.Get("user").(*jwt.Token)
	// claims := token.Claims.(jwt.MapClaims)
	// role, ok := claims["role"].(string)
	// if !ok || role != "admin" {
	// 	return echo.ErrUnauthorized
	// }

	tool := model.Tool{}
	c.Bind(&tool)

	err := config.DB.Save(&tool).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to add tool",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully added the tool",
		"tool":    tool,
	})
}

// Get_tools
func Get_tools(c echo.Context) error {
	var tools []model.Tool

	err := config.DB.Find(&tools).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all tools",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully to get all the tools",
		"tools":   tools,
	})
}

// Modify_stock
func Modify_stock(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var tool model.Tool

	err_cek1 := config.DB.First(&tool, id).Error
	if err_cek1 != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "tool id not found",
			"error":   err_cek1.Error(),
		})
	}

	var stock model.Stock
	err_cek2 := c.Bind(&stock)
	if err_cek2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind",
			"error":   err_cek2.Error(),
		})
	}

	tool.Stock += stock.Stock
	config.DB.Save(&tool)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully added stock",
		"tool":    tool,
	})
}

// Delete_tool
func Delete_tool(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var tool model.Tool

	err := config.DB.First(&tool, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "tool id not found",
			"error":   err.Error(),
		})
	}
	config.DB.Delete(&tool)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully delete tool",
		"tool":    tool,
	})
}
