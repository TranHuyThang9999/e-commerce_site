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

type UserCollection struct {
	collection *gorm.DB
}

func NewAccountRepository(cf *configs.Configs, user *adapter.PostGresql) ports.RepositoryAccount {
	return &UserCollection{
		collection: user.CreateCollection(),
	}
}

func (u *UserCollection) CreateAccount(ctx context.Context, tx *gorm.DB, req *model.Account) error {
	result := tx.Create(req)
	return result.Error
}
func (u *UserCollection) FindByFormAccount(ctx context.Context, req *model.AccountReqFindByForm) ([]*model.Account, error) {
	var accounts []*model.Account
	result := u.collection.Where(&model.Account{
		ID:          req.ID,
		FirstName:   req.FirstName,
		LastName:    req.FirstName,
		Age:         req.Age,
		Address:     req.Address,
		Gender:      req.Gender,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		UserName:    req.UserName,
		StoreName:   req.StoreName,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.CreatedAt,
	}).Find(&accounts)
	return accounts, result.Error
}
func (u *UserCollection) UpdateAccount(ctx context.Context, req *model.Account) error {
	result := u.collection.Where("id = ?", req.ID).Updates(req)
	return result.Error
}
func (u *UserCollection) GetInfomationByEmail(ctx context.Context, email string) (*model.Account, error) {
	var account *model.Account
	result := u.collection.Where("email = ?", email).First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return account, result.Error
}
func (u *UserCollection) GetInfomationByUserName(ctx context.Context, userName string) (*model.Account, error) {
	var account *model.Account
	result := u.collection.Where("user_name = ?", userName).First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return account, result.Error
}
func (u *UserCollection) GetInfomationByPhoneumber(ctx context.Context, phone_number string) (*model.Account, error) {
	var account *model.Account
	result := u.collection.Where("phone_number = ?", phone_number).First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return account, result.Error
}
func (u *UserCollection) GetInfomationByStoreName(ctx context.Context, store_name string) (*model.Account, error) {
	var account *model.Account
	result := u.collection.Where("store_name = ?", store_name).First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return account, result.Error
}
