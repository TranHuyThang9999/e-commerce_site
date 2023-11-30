package usecases

import (
	"context"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/common/log"
	"ecommerce_site/src/core/enums"
	"ecommerce_site/src/core/ports"
	"strconv"
)

type ImageStorageUseCase struct {
	file ports.RepositoryUploadImage
}

func NewFileUseCase(
	file ports.RepositoryUploadImage,
) *ImageStorageUseCase {
	return &ImageStorageUseCase{
		file: file,
	}
}

func (u *ImageStorageUseCase) DeleteImageById(ctx context.Context, id string) (*model.DeleteImageByIdResp, error) {

	numOffset, err := strconv.Atoi(id)
	if err != nil {
		return &model.DeleteImageByIdResp{
			Result: model.Result{
				Code:    enums.CONVERT_TO_NUMBER_CODE,
				Message: enums.CONVERT_TO_NUMBER_MESS,
			},
		}, nil
	}

	infoImage, err := u.file.GetFileById(ctx, int64(numOffset))
	if err != nil {
		return &model.DeleteImageByIdResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if infoImage == nil {
		return &model.DeleteImageByIdResp{
			Result: model.Result{
				Code:    enums.IMAGE_NOT_FOUND_CODE,
				Message: enums.IMAGE_NOT_FOUND_MESS,
			},
		}, nil
	}

	err = u.file.DeleteImageById(ctx, int64(numOffset))

	if err != nil {
		return &model.DeleteImageByIdResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	return &model.DeleteImageByIdResp{
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (u *ImageStorageUseCase) GetAllImageForUserNameByIdProduct(ctx context.Context, idProduct string) (*model.ImagesRespGetListByIdProduct, error) {

	numOffset, err := strconv.ParseInt(idProduct, 10, 64)
	log.Info(idProduct)
	if err != nil {
		return &model.ImagesRespGetListByIdProduct{
			Result: model.Result{
				Code:    enums.CONVERT_TO_NUMBER_CODE,
				Message: enums.CONVERT_TO_NUMBER_MESS,
			},
		}, nil
	}

	images, err := u.file.GetAllImageForUserNameByIdProduct(ctx, int64(numOffset))
	if err != nil {
		return &model.ImagesRespGetListByIdProduct{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(images) == 0 {
		return &model.ImagesRespGetListByIdProduct{
			Result: model.Result{
				Code:    enums.IMAGES_EMPTY_CODE,
				Message: enums.IMAGEST_EMPTY_MESS,
			},
		}, nil
	}
	return &model.ImagesRespGetListByIdProduct{
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		ImageStorage: images,
	}, nil

}
