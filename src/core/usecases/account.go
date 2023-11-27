package usecases

import (
	"context"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/common/imgbb"
	"ecommerce_site/src/common/utils"
	"ecommerce_site/src/core/enums"
	"ecommerce_site/src/core/ports"
)

type UseCaseAccount struct {
	account     ports.RepositoryAccount
	transaction ports.RepositoryTransaction
	roler       ports.RepositoryRole
}

func NewUseCaseAccount(
	account ports.RepositoryAccount,
	transaction ports.RepositoryTransaction,
	roler ports.RepositoryRole,
) *UseCaseAccount {
	return &UseCaseAccount{
		account:     account,
		transaction: transaction,
		roler:       roler,
	}
}
func (u *UseCaseAccount) CreateAccount(ctx context.Context, req *model.AccountReqCreate) (*model.AccountRespCreate, error) {

	checkEmailExits, err := u.account.GetInfomationByEmail(ctx, req.Email)
	if err != nil {
		return &model.AccountRespCreate{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if checkEmailExits != nil {
		return &model.AccountRespCreate{
			Result: model.Result{
				Code:    enums.EMAIL_ACCOUNT_EXITS_CODE,
				Message: enums.EMAIL_ACCOUNT_EXITS_MESS,
			},
		}, nil
	}

	checkUserName, err := u.account.GetInfomationByUserName(ctx, req.UserName)
	if err != nil {
		return &model.AccountRespCreate{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if checkUserName != nil {
		return &model.AccountRespCreate{
			Result: model.Result{
				Code:    enums.ACCOUNT_EXIST_CODE,
				Message: enums.ACCOUNT_EXIST_MESS,
			},
		}, nil
	}

	checkPhoneNumber, err := u.account.GetInfomationByPhoneumber(ctx, req.PhoneNumber)
	if err != nil {
		return &model.AccountRespCreate{
			Id: 0,
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if checkPhoneNumber != nil {
		return &model.AccountRespCreate{
			Result: model.Result{
				Code:    enums.PHONE_NUMBER_EXITS_CODE,
				Message: enums.PHONE_NUMBER_EXITS_MESS,
			},
		}, nil
	}

	// checkStore, err := u.account.GetInfomationByStoreName(ctx, req.StoreName)
	// if err != nil {
	// 	return &model.AccountRespCreate{
	// 		Id: 0,
	// 		Result: model.Result{
	// 			Code:    enums.DB_ERR_CODE,
	// 			Message: enums.DB_ERR_MESS,
	// 		},
	// 	}, nil
	// }
	// if checkStore != nil {
	// 	return &model.AccountRespCreate{
	// 		Result: model.Result{
	// 			Code:    enums.STORE_NAME_EXITS_CODE,
	// 			Message: enums.STORE_NAME_EXITS_MESS,
	// 		},
	// 	}, nil
	// }

	idUser := utils.GenerateUniqueUUid()
	idRole := utils.GenerateUniqueUUid()
	newCode := utils.CodeOPT()
	tx, err := u.transaction.BeginTransaction(ctx)
	utlAvatar := ""
	if err != nil {
		return &model.AccountRespCreate{
			Result: model.Result{
				Code:    enums.TRANSACTION_INVALID_CODE,
				Message: enums.TRANSACTION_INVALID_MESS,
			},
		}, nil
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return &model.AccountRespCreate{
			Result: model.Result{
				Code:    enums.HASH_PASSWORD_ERR,
				Message: enums.HASH_PASSWORD_ERR_MESS,
			},
		}, nil
	}

	if req.File != nil {
		url, err := imgbb.ProcessImageSign(req.File)
		if err != nil {
			return &model.AccountRespCreate{
				Result: model.Result{
					Code:    enums.ERROR_SAVE_IMAGE_CODE,
					Message: enums.ERRORL_SAVE_IMAGE_MESS,
				},
			}, nil
		}
		utlAvatar = url.URL

	} else {
		utlAvatar = ""
	}
	err = utils.SendEmail(req.Email, int64(newCode))
	if err != nil {
		tx.Rollback()
		return &model.AccountRespCreate{
			Result: model.Result{
				Code:    enums.SEND_EMAIL_CODE_ERROR,
				Message: enums.SEND_EMAIL_MESS_ERROR,
			},
		}, nil
	}
	err = u.account.CreateAccount(ctx, tx, &model.Account{
		ID:          idUser,
		IDRole:      idRole,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Age:         req.Age,
		Address:     req.Address,
		Gender:      req.Gender,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		UserName:    req.UserName,
		Password:    hashedPassword,
		OTPCode:     int64(newCode),
		OTPExpiry:   0,
		IsVerified:  enums.NOT_VERIFIED,
		Notes:       req.Notes,
		// StoreName:   req.StoreName,
		CreatedAt:   int(utils.GetCurrentTimestamp()),
		UpdatedAt:   int(utils.GetCurrentTimestamp()),
		Avatar:      utlAvatar,
		DateOfBirth: req.DateOfBirth,
	})
	if err != nil {
		tx.Rollback()
		return &model.AccountRespCreate{
			Id: 0,
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	err = u.roler.AddRole(ctx, tx, &model.Role{
		ID:     idRole,
		Admin:  enums.ROLE_NOT_ADMIN,
		Seller: enums.ROLE_USER_BUYER_ACTIVE,
		Buyer:  enums.ROLE_USER_SELLER_LOCK,
	})
	if err != nil {
		tx.Rollback()
		return &model.AccountRespCreate{
			Id: 0,
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	tx.Commit()
	return &model.AccountRespCreate{
		Id: 0,
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
