package usecases

import (
	"context"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/common/dto"
	"ecommerce_site/src/common/imgbb"
	"ecommerce_site/src/common/log"
	"ecommerce_site/src/common/utils"
	"ecommerce_site/src/core/enums"
	"ecommerce_site/src/core/ports"
	"strconv"
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
	//var list_id_image = make([]int64, 0)

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

	inforImages, err := imgbb.ProcessImages(req.Files)
	if err != nil {
		tx.Rollback()
		return &model.ProductRespCreate{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	//
	if len(inforImages) > 0 {
		for _, file := range inforImages {
			id_image := utils.GenerateUniqueUUid()
			listInfodata = append(listInfodata, &model.ImageStorage{
				ID:        id_image,
				Url:       file.URL,
				IDUser:    account.ID,
				IDProduct: idProduct,
			})
		}

		err = u.file.UploadImageMutileFile(ctx, tx, listInfodata) //
		if err != nil {
			log.Infof("data ", err) //
			tx.Rollback()
			return &model.ProductRespCreate{
				Result: model.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, nil
		}
	}

	err = u.product.AddProduct(ctx, tx, &model.Product{
		ID:             idProduct,
		IDUser:         account.ID,
		NameProduct:    req.NameProduct,
		Quantity:       req.Quantity,
		SellStatus:     req.SellStatus,
		Price:          req.Price,
		Discount:       req.Discount,
		Manufacturer:   req.Manufacturer,
		CreatedAt:      int(utils.GetCurrentTimestamp()),
		UpdatedAt:      int(utils.GetCurrentTimestamp()),
		Describe:       req.Describe,
		IDTypeProduct:  req.IDTypeProduct,
		NumberOfPhotos: len(inforImages),
		//ListIdImage:   list_id_image_str,
	})
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

func (u *ProductUseCase) GetListProductUserSeller(ctx context.Context, req *dto.ProductReqFindByForm) (*model.ProductListRespSeller, error) {

	var limit int
	var offset int

	account, err := u.user.GetInfomationByUserName(ctx, req.UserName)
	if err != nil {
		return &model.ProductListRespSeller{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if account == nil {
		return &model.ProductListRespSeller{
			Result: model.Result{
				Code:    enums.ACCOUNT_NOT_EXIST_CODE,
				Message: enums.ACCOUNT_NOT_EXIST_MESS,
			},
		}, nil
	}

	if req.Limit == 0 {
		limit = 10
	} else {
		limit = req.Limit
	}

	offset = (req.Offset - 1) * limit

	productsById, err := u.product.FindByForm(ctx, &model.ProductReqFindByForm{
		ID:            req.ID,
		IdUser:        account.ID,
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
	}, offset, limit)

	if err != nil {
		return &model.ProductListRespSeller{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(productsById) == 0 {
		return &model.ProductListRespSeller{
			Total: 0,
			Result: model.Result{
				Code:    enums.PRODUCT_EMPTY_CODE,
				Message: enums.PRODUCT_EMPTY_MESS,
			},
		}, nil
	}

	return &model.ProductListRespSeller{
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Total:   len(productsById),
		Product: productsById,
	}, nil
}
func (u *ProductUseCase) DeleteProductById(ctx context.Context, idProduct string) (*model.ProductDeleteByIdResp, error) {

	idNumber, err := strconv.ParseInt(idProduct, 10, 64)
	if err != nil {
		return &model.ProductDeleteByIdResp{
			Result: model.Result{
				Code:    enums.CONVERT_TO_NUMBER_CODE,
				Message: enums.CONVERT_TO_NUMBER_MESS,
			},
		}, nil
	}

	tx, err := u.trans.BeginTransaction(ctx)
	if err != nil {
		return &model.ProductDeleteByIdResp{
			Result: model.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, err
	}

	err = u.product.DeleteProductById(ctx, tx, idNumber)
	if err != nil {
		tx.Rollback()
		return &model.ProductDeleteByIdResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = u.file.DeleteImagesByIdProduct(ctx, tx, idNumber)
	if err != nil {
		tx.Rollback()
		return &model.ProductDeleteByIdResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &model.ProductDeleteByIdResp{
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (u *ProductUseCase) UpdateProductById(ctx context.Context, req *model.ProductUpdateByIdReq) (*model.ProductUpdateByIdResp, error) {

	updateAt := utils.GetCurrentTimestamp()
	var listInfodata []*model.ImageStorage

	tx, err := u.trans.BeginTransaction(ctx)
	if err != nil {
		return &model.ProductUpdateByIdResp{
			Result: model.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, err
	}

	productResp, err := u.product.FindProductById(ctx, req.ID)
	if err != nil {
		return &model.ProductUpdateByIdResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	inforImages, err := imgbb.ProcessImages(req.Files)
	if err != nil {
		tx.Rollback()
		return &model.ProductUpdateByIdResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if len(inforImages) > 0 {
		for _, file := range inforImages {
			id_image := utils.GenerateUniqueUUid()
			listInfodata = append(listInfodata, &model.ImageStorage{
				ID:        id_image,
				Url:       file.URL,
				IDUser:    req.IDUser,
				IDProduct: req.ID,
			})
		}

		err = u.file.UploadImageMutileFile(ctx, tx, listInfodata) //
		if err != nil {
			log.Infof("data ", err) //
			tx.Rollback()
			return &model.ProductUpdateByIdResp{
				Result: model.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, nil
		}
	}

	err = u.product.UpdateProductById(ctx, tx, &model.Product{
		ID:             req.ID,
		IDUser:         req.IDUser,
		NameProduct:    req.Describe,
		Quantity:       req.Quantity,
		SellStatus:     req.SellStatus,
		Price:          req.Price,
		Discount:       req.Discount,
		Manufacturer:   req.Manufacturer,
		CreatedAt:      productResp.CreatedAt,
		UpdatedAt:      int(updateAt),
		Describe:       req.Describe,
		IDTypeProduct:  req.IDTypeProduct,
		NumberOfPhotos: len(listInfodata),
	})
	if err != nil {
		tx.Rollback()
		return &model.ProductUpdateByIdResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &model.ProductUpdateByIdResp{
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
