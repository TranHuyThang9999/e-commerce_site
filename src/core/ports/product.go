package ports

import (
	"context"
	"ecommerce_site/src/adapter/model"

	"gorm.io/gorm"
)

type RepositoryProducts interface {
	AddProduct(ctx context.Context, tx *gorm.DB, req *model.Product) error
	FindByForm(ctx context.Context, req *model.ProductReqFindByForm, offset, limit int) ([]*model.Product, error)
	DeleteProductById(ctx context.Context, tx *gorm.DB, id int64) error
	UpdateProductById(ctx context.Context, tx *gorm.DB, req *model.Product) error
	FindProductById(ctx context.Context, id int64) (*model.Product, error)
}
