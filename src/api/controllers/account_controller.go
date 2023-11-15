package controllers

import (
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/core/usercases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	*baseController
	account *usercases.UseCaseAccount
}

func NewControllerAccount(account *usercases.UseCaseAccount) *AccountController {
	return &AccountController{
		account: account,
	}
}
func (u *AccountController) CreateAccount(ctx *gin.Context) {

	var req model.AccountReqCreate

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	file, err := ctx.FormFile("image")
	if err != nil && err != http.ErrMissingFile && err != http.ErrNotMultipart {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Không thể tải ảnh lên.",
		})
		return
	}

	if err := u.validateRequest(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	req.File = file
	resp, err := u.account.CreateAccount(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error 1 ": resp})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
