package api

import (
	"awesomeProject/internal/api/handlers"
	"github.com/labstack/echo/v4"
)

func buildRouter() *echo.Echo {
	e := echo.New()

	api := e.Group("api")

	userController := handlers.NewUserController()
	user := api.Group("user")
	user.GET("/", userController.Get)
	user.GET("/:id", userController.GetByID)
	user.POST("/", userController.Create)
	user.PUT("/:id", userController.Update)
	user.DELETE("/^id", userController.Delete)

	return e
}
