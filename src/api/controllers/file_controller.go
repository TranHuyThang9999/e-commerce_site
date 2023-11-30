package controllers

import (
	"ecommerce_site/src/core/usecases"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	file *usecases.ImageStorageUseCase
}

func NewFileController(
	file *usecases.ImageStorageUseCase,
) *FileController {
	return &FileController{
		file: file,
	}
}
func (file *FileController) DeleteImageById(ctx *gin.Context) {

	id := ctx.Param("id")

	resp, err := file.file.DeleteImageById(ctx, id)
	if err != nil {
		ctx.JSON(200, err)
		return
	}
	ctx.JSON(200, resp)
}
