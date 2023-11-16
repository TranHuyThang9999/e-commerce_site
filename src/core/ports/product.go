package ports

import (
	"context"
	"ecommerce_site/src/adapter/model"
)

type RepositoryProducts interface {
	AddProduct(ctx context.Context, req *model.Product) error
	FindByForm(ctx context.Context, req *model.ProductReqFindByForm, limit, offset int) ([]*model.ProductRespFindByForm, error)
}
