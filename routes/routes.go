package routes

import (
	"github.com/labstack/echo"
	"github.com/sigitprd/simple-crud-rest/controllers"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "hello world")
	})
	e.GET("/user", controllers.FetchUser)
	e.GET("/user:id", controllers.FetchUser)
	e.POST("/user", controllers.StoreUser)
	e.PUT("/user:id", controllers.UpdateUser)
	e.DELETE("/user:id", controllers.DeleteUser)

	return e
}
