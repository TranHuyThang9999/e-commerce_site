package postgresql

import (
	"context"
	"ecommerce_site/src/adapter"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/configs"
	"ecommerce_site/src/core/ports"

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
func (u *ProductCollection) FindByForm(ctx context.Context, req *model.ProductReqFindByForm, offset, limit int) ([]*model.ProductRespFindByForm, error) {
	var products []*model.ProductRespFindByForm
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
	}).Offset(offset).Limit(limit).Order("created_at desc").Model(&model.Product{}).Find(&products)
	return products, result.Error
}
