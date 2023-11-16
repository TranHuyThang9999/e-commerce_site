package usecases

import (
	"context"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/common/imgbb"
	"ecommerce_site/src/common/utils"
	"ecommerce_site/src/core/enums"
	"ecommerce_site/src/core/ports"
)

type ProductUseCase struct {
	product ports.RepositoryProducts
	user    ports.RepositoryAccount
	file    ports.RepositoryUploadImage
	trans   ports.RepositoryTransaction
}

func NewProductUseCase(
	product ports.RepositoryProducts,
	user ports.RepositoryAccount,
	file ports.RepositoryUploadImage,
	trans ports.RepositoryTransaction,
) *ProductUseCase {
	return &ProductUseCase{
		product: product,
		user:    user,
		file:    file,
		trans:   trans,
	}
}

func (u *ProductUseCase) AddProduct(ctx context.Context, req *model.ProductReqCreate) (*model.ProductRespCreate, error) {

	idProduct := utils.GenerateUniqueUUid()
	var listInfodata []*model.ImageStorage

	tx, err := u.trans.BeginTransaction(ctx)
	if err != nil {
		return &model.ProductRespCreate{
			Result: model.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, err
	}

	account, err := u.user.GetInfomationByUserName(ctx, req.UserName)
	if err != nil {
		return &model.ProductRespCreate{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if account == nil {
		return &model.ProductRespCreate{
			Result: model.Result{
				Code:    enums.ACCOUNT_NOT_EXIST_CODE,
				Message: enums.ACCOUNT_NOT_EXIST_MESS,
			},
		}, nil
	}

	err = u.product.AddProduct(ctx, &model.Product{
		ID:            idProduct,
		IDUser:        account.ID,
		NameProduct:   req.NameProduct,
		Quantity:      req.Quantity,
		SellStatus:    req.SellStatus,
		Price:         req.Price,
		Discount:      req.Discount,
		Manufacturer:  req.Manufacturer,
		CreatedAt:     int(utils.GetCurrentTimestamp()),
		UpdatedAt:     int(utils.GetCurrentTimestamp()),
		Describe:      req.Describe,
		IDTypeProduct: req.IDTypeProduct,
	})
	if err != nil {
		return &model.ProductRespCreate{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	inforImages, err := imgbb.ProcessImages(req.Files)
	if err != nil {
		tx.Rollback()
		return &model.ProductRespCreate{
			Result: model.Result{
				Code:    1,
				Message: "err 1",
			},
		}, nil
	}

	for _, file := range inforImages {
		listInfodata = append(listInfodata, &model.ImageStorage{
			ID:        utils.GenerateUniqueUUid(),
			Url:       file.URL,
			IDUser:    account.ID,
			IDProduct: idProduct,
		})
	}
	err = u.file.UploadImageMutileFile(ctx, tx, listInfodata)
	if err != nil {
		tx.Rollback()
		return &model.ProductRespCreate{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &model.ProductRespCreate{
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil

}