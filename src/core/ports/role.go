package ports

import (
	"context"
	"ecommerce_site/src/adapter/model"

	"gorm.io/gorm"
)

type RepositoryRole interface {
	AddRole(ctx context.Context, tx *gorm.DB, req *model.Role) error
}
