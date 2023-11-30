package postgresql

import (
	"context"
	"ecommerce_site/src/adapter"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/configs"
	"ecommerce_site/src/core/ports"
	"errors"

	"gorm.io/gorm"
)

type ProductCollection struct {
	collection *gorm.DB
}

func NewProductRepository(cf *configs.Configs, product *adapter.PostGresql) ports.RepositoryProducts {
	return &ProductCollection{
		collection: product.CreateCollection(),
	}
}
func (u *ProductCollection) AddProduct(ctx context.Context, tx *gorm.DB, req *model.Product) error {
	result := tx.Create(req)
	return result.Error
}
func (u *ProductCollection) FindByForm(ctx context.Context, req *model.ProductReqFindByForm, offset, limit int) ([]*model.Product, error) {
	var products []*model.Product
	result := u.collection.Where(&model.Product{
		ID:            req.ID,
		IDUser:        req.IdUser,
		NameProduct:   req.Describe,
		Quantity:      req.Quantity,
		SellStatus:    req.SellStatus,
		Price:         req.Price,
		Discount:      req.Discount,
		Manufacturer:  req.Manufacturer,
		CreatedAt:     req.CreatedAt,
		UpdatedAt:     req.UpdatedAt,
		Describe:      req.Describe,
		IDTypeProduct: req.IDTypeProduct,
	}).Offset(offset).Limit(limit).Order("created_at desc").Find(&products)
	return products, result.Error
}
func (u *ProductCollection) DeleteProductById(ctx context.Context, tx *gorm.DB, id int64) error {
	result := tx.Where("id = ?", id).Delete(&model.Product{})
	return result.Error
}
func (u *ProductCollection) UpdateProductById(ctx context.Context, tx *gorm.DB, req *model.Product) error {
	result := tx.Save(req)
	return result.Error
}
func (u *ProductCollection) FindProductById(ctx context.Context, id int64) (*model.Product, error) {
	var product *model.Product
	result := u.collection.Where("id = ?", id).First(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return product, result.Error

}
