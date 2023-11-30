package controllers

import (
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/common/dto"
	"ecommerce_site/src/common/imgbb"
	"ecommerce_site/src/core/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerProduct struct {
	*baseController
	ctl *usecases.ProductUseCase
}

func NewControllerProduct(ctl *usecases.ProductUseCase) *ControllerProduct {
	return &ControllerProduct{
		ctl: ctl,
	}
}

func (t *ControllerProduct) AddProduct(c *gin.Context) {

	var req model.ProductReqCreate

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	files, err := imgbb.GetUploadedFiles(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Files = files

	resp, err := t.ctl.AddProduct(c, &req)
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, resp)
}
func (t *ControllerProduct) GetListProduct(c *gin.Context) {

	var req dto.ProductReqFindByForm

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := t.validateRequest(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	resp, err := t.ctl.GetListProductUserSeller(c, &req)
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, resp)

}
func (t *ControllerProduct) DeleteProductById(c *gin.Context) {

	id := c.Query("id")

	resp, err := t.ctl.DeleteProductById(c, id)
	if err != nil {
		c.JSON(200, err)
		return
	}
	c.JSON(200, resp)
}
