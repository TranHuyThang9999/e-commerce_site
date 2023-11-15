package ports

import (
	"context"
	"ecommerce_site/src/adapter/model"

	"gorm.io/gorm"
)

type RepositoryAccount interface {
	CreateAccount(ctx context.Context, tx *gorm.DB, req *model.Account) error
	FindByFormAccount(ctx context.Context, req *model.AccountReqFindByForm) ([]*model.Account, error)
	UpdateAccount(ctx context.Context, req *model.Account) error
	GetInfomationByEmail(ctx context.Context, email string) (*model.Account, error)
	GetInfomationByUserName(ctx context.Context, userName string) (*model.Account, error)
	GetInfomationByPhoneumber(ctx context.Context, phonenUmber string) (*model.Account, error)
	GetInfomationByStoreName(ctx context.Context, storeName string) (*model.Account, error)
}
