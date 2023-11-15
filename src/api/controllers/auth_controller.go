package controllers

import (
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/core/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auth *usecases.JwtUseCase
}

func NewAuthController(auth *usecases.JwtUseCase) *AuthController {
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

func (u *AuthController) VerifiedAccount(ctx *gin.Context) {
	userName := ctx.Query("userName")
	code := ctx.Query("code")

	resp, err := u.auth.VerifiedAccount(ctx, userName, code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 1 ": err})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
