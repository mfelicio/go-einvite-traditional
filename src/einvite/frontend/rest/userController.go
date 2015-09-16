package rest

import (
	"einvite/common/contracts"
	"einvite/common/services"
	"einvite/framework"
)

type UserController struct {
	userService services.UserService
}

func (this UserController) Who(ctx framework.WebContext) framework.WebResult {

	return ctx.Text("Anonymous")
}

func (this UserController) CreateUser(ctx framework.WebContext) framework.WebResult {

	return ctx.Forbidden("Thou shall not create an user")
}

func (this UserController) ListUsers(ctx framework.WebContext) framework.WebResult {

	users, err := this.userService.List()

	if err == nil {
		return ctx.Json(users)
	}

	return ctx.Error(err)
}

func (this UserController) GetUser(ctx framework.WebContext) framework.WebResult {

	name, _ := ctx.Param("name")
	email, _ := ctx.Param("email")

	user, err := this.userService.Create(&contracts.User{Email: email, Name: name})

	if err == nil {
		return ctx.Json(user)
	}

	return ctx.Error(err)
}

func NewUserController(userService services.UserService) *UserController {

	return &UserController{userService}
}
