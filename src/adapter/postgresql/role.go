package postgresql

import (
	"context"
	"ecommerce_site/src/adapter"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/configs"
	"ecommerce_site/src/core/ports"

	"gorm.io/gorm"
)

type roleCollection struct {
	collection *gorm.DB
}

func NewRoleRepository(cf *configs.Configs, user *adapter.PostGresql) ports.RepositoryRole {
	return &roleCollection{
		collection: user.CreateCollection(),
	}
}

func (u *roleCollection) AddRole(ctx context.Context, tx *gorm.DB, req *model.Role) error {
	result := tx.Create(req)
	return result.Error
}
