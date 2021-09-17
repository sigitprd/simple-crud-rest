package controllers

import (
	"github.com/labstack/echo"
	"github.com/sigitprd/simple-crud-rest/models"
	"net/http"
)

func FetchUser(ctx echo.Context) error {
	result, err := models.FetchUser(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, result)
}
