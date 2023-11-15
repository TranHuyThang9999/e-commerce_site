package controllers

import (
	"ecommerce_site/src/core/entities"
	"ecommerce_site/src/core/usercases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllersUser struct {
	ctl *usercases.UserUseCase
}

func NewControllerUser(ctl *usercases.UserUseCase) *ControllersUser {
	return &ControllersUser{
		ctl: ctl,
	}
}
func (t *ControllersUser) AddProfile(c *gin.Context) {

	var req entities.UsersReq

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := t.ctl.AddProfile(c, &req)
	if err != nil {
		c.JSON(200, resp)
		return
	}
	c.JSON(200, resp)
}
