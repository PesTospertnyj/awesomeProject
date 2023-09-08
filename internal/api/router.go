package api

import (
	"github.com/awesomeProject/internal/api/handlers"
	"github.com/awesomeProject/internal/services"
	"github.com/awesomeProject/internal/storage/repositories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Router struct {
	userController handlers.UserController
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{
		userController: handlers.NewUserController(
			services.NewUserService(
				repositories.NewUserRepository(db),
			),
		),
	}
}

func (r *Router) build() *echo.Echo {
	e := echo.New()

	api := e.Group("api")

	user := api.Group("user")
	user.GET("/", r.userController.Get)
	user.GET("/:id", r.userController.GetByID)
	user.POST("/", r.userController.Create)
	user.PUT("/:id", r.userController.Update)
	user.DELETE("/:id", r.userController.Delete)

	return e
}
