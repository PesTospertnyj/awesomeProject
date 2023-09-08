package handlers

import (
	"github.com/awesomeProject/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	service services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		service: userService,
	}
}

func (uc *UserController) Get(c echo.Context) error {
	return c.String(http.StatusOK, "get all users")
}

func (uc *UserController) GetByID(c echo.Context) error {
	return c.String(http.StatusOK, "get user by id")
}

func (uc *UserController) Create(c echo.Context) error {
	return c.String(http.StatusOK, "create users")
}

func (uc *UserController) Update(c echo.Context) error {
	return c.String(http.StatusOK, "update users")
}

func (uc *UserController) Delete(c echo.Context) error {
	return c.String(http.StatusOK, "delete users")
}
