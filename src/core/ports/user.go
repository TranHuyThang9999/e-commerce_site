package ports

import (
	"context"
	"ecommerce_site/src/adapter/model"
)

type RepositoryUser interface {
	AddProfile(ctx context.Context, req *model.Users) error
}
