package postgresql

import (
	"context"
	"ecommerce_site/src/adapter"
	"ecommerce_site/src/adapter/configs"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/core/ports"

	"gorm.io/gorm"
)

type CollectionUser struct {
	user *gorm.DB
}

func NewRepositoryUser(cf *configs.Configs, user *adapter.PostGresql) ports.RepositoryUser {
	return &CollectionUser{
		user: user.CreateCollection(),
	}
}
func (user *CollectionUser) AddProfile(ctx context.Context, req *model.Users) error {
	result := user.user.Create(req)
	return result.Error
}
