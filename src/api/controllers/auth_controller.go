package controllers

import (
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/core/usercases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auth *usercases.JwtUseCase
}

func NewAuthController(auth *usercases.JwtUseCase) *AuthController {
	return &AuthController{
		auth: auth,
	}
}
func (u *AuthController) Login(ctx *gin.Context) {

	var req model.LoginReq

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	resp, err := u.auth.LoginAccount(ctx, req.UserName, req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 1 ": err})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
