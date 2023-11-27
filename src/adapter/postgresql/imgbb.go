package postgresql

import (
	"context"
	"ecommerce_site/src/adapter"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/configs"
	"ecommerce_site/src/core/ports"

	"gorm.io/gorm"
)

type CollectionUploadFile struct {
	collection *gorm.DB
}

func NewUploadFileRepository(cf *configs.Configs, img *adapter.PostGresql) ports.RepositoryUploadImage {
	return &CollectionUploadFile{
		collection: img.CreateCollection(),
	}
}

func (u *CollectionUploadFile) UploadImageSingleFile(ctx context.Context, tx *gorm.DB, req *model.ImageStorage) error {
	result := tx.Create(req)
	return result.Error
}
func (u *CollectionUploadFile) GetAllImageForUserNameByIdProduct(ctx context.Context, idProduct int64) ([]*model.ImageStorage, error) {
	var desImageProduct []*model.ImageStorage
	result := u.collection.Where("id_product = ?", idProduct).Find(&desImageProduct)
	return desImageProduct, result.Error
}
func (u *CollectionUploadFile) UpdateImageByIdProduct(ctx context.Context, req *model.ImageStorage) error {
	result := u.collection.Model(&model.ImageStorage{}).Updates(req)
	return result.Error
}
func (u *CollectionUploadFile) UploadImageMutileFile(ctx context.Context, tx *gorm.DB, req []*model.ImageStorage) error {
	result := tx.Create(req)
	return result.Error
}
func (u *CollectionUploadFile) DeleteImageById(ctx context.Context, IdImage int64) error {
	result := u.collection.Where("id = ?", IdImage).Delete(&model.ImageStorage{})
	return result.Error
}
func (u *CollectionUploadFile) FindBymultipleId(ctx context.Context, ids []int64) ([]*model.ImageStorage, error) {
	var desImageProduct []*model.ImageStorage
	result := u.collection.Where("id IN (?)", ids).Find(&desImageProduct)

	return desImageProduct, result.Error
}
func (u *CollectionUploadFile) FindAllImages(ctx context.Context) ([]*model.ImageStorage, error) {
	var desImageProduct []*model.ImageStorage
	result := u.collection.Find(&desImageProduct)
	return desImageProduct, result.Error
}
