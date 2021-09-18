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

func StoreUser(ctx echo.Context) error {
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	roles := ctx.FormValue("roles")

	result, err := models.StoreUser(username, email, roles)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, result)
}

func UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	roles := ctx.FormValue("roles")

	result, err := models.UpdateUser(id, username, email, roles)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, result)
}

func DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")

	result, err := models.DeleteUser(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, result)
}
