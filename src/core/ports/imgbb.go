package ports

import (
	"context"
	"ecommerce_site/src/adapter/model"

	"gorm.io/gorm"
)

type RepositoryUploadImage interface {
	UploadImageSingleFile(ctx context.Context, tx *gorm.DB, req *model.ImageStorage) error
	GetAllImageForUserNameByIdProduct(ctx context.Context, idProduct int64) ([]*model.ImageStorage, error)
	UpdateImageByIdProduct(ctx context.Context, req *model.ImageStorage) error // chua dung
	UploadImageMutileFile(ctx context.Context, tx *gorm.DB, req []*model.ImageStorage) error
	DeleteImageById(ctx context.Context, IdImage int64) error
	FindBymultipleId(ctx context.Context, ids []int64) ([]*model.ImageStorage, error)
}
