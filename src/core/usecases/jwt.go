package usecases

import (
	"context"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/common/log"
	"ecommerce_site/src/common/utils"
	"ecommerce_site/src/configs"
	"ecommerce_site/src/core/enums"
	"ecommerce_site/src/core/ports"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtUseCase struct {
	config             *configs.Configs
	expAccessToken     time.Duration
	expRefreshToken    time.Duration
	userRepositoryPort ports.RepositoryAccount
}

func NewJwtUseCase(cf *configs.Configs, userRepositoryPort ports.RepositoryAccount) (*JwtUseCase, error) {
	expAccessToken, err := time.ParseDuration(cf.ExpireAccess)
	if err != nil {
		return nil, fmt.Errorf("expire access token has wrong format: %s", err)
	}
	expRefreshToken, err := time.ParseDuration(cf.ExpireRefresh)
	if err != nil {
		return nil, fmt.Errorf("expire refresh token has wrong format: %s", err)
	}
	return &JwtUseCase{
		config:             cf,
		expAccessToken:     expAccessToken,
		expRefreshToken:    expRefreshToken,
		userRepositoryPort: userRepositoryPort,
	}, nil
}

func (u *JwtUseCase) encrypt(secret string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (u *JwtUseCase) Decrypt(tokenString string) (*model.UserJwtClaim, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&model.UserJwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(u.config.AccessSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.UserJwtClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, nil
}

func (u *JwtUseCase) generateToken(id int64, userName string, idRole int64) (*model.JwtToken, error) {
	userClaim := func(expire time.Duration) *model.UserJwtClaim {
		return &model.UserJwtClaim{
			Id:       id,
			UserName: userName,
			Role:     idRole,
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: time.Now().Add(expire).Unix(),
			},
		}
	}

	accessToken, err := u.encrypt(u.config.AccessSecret, userClaim(u.expAccessToken))
	if err != nil {
		log.Error(err, "Error when generating access token")
		return nil, err
	}
	refreshToken, err := u.encrypt(u.config.RefreshSecret, userClaim(u.expAccessToken))

	if err != nil {
		log.Error(err, "Error when generating refresh token")
		return nil, err
	}

	return &model.JwtToken{
		AccessToken:  accessToken,
		AtExpires:    int64(u.expAccessToken / time.Second),
		RefreshToken: refreshToken,
		RtExpires:    int64(u.expRefreshToken / time.Second),
	}, nil
}
func (u *JwtUseCase) LoginAccount(ctx context.Context, userName, passWord string) (*model.ResponseLogin, error) {

	account, err := u.userRepositoryPort.GetInfomationByUserName(ctx, userName)
	if err != nil {
		return &model.ResponseLogin{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	if account == nil {
		return &model.ResponseLogin{
			Result: model.Result{
				Code:    enums.ACCOUNT_OR_PASSWORD_WRONG_CODE,
				Message: enums.ACCOUNT_OR_PASSWORD_WRONG_MESS,
			},
		}, nil
	}

	err = utils.ComparePassword(account.Password, passWord)
	if err != nil {
		return &model.ResponseLogin{
			Result: model.Result{
				Code:    enums.ACCOUNT_OR_PASSWORD_WRONG_CODE,
				Message: enums.ACCOUNT_OR_PASSWORD_WRONG_MESS,
			},
		}, nil
	}

	if account.IsVerified == enums.NOT_VERIFIED {
		return &model.ResponseLogin{
			Result: model.Result{
				Code:    enums.ACCOUNT_NOT_VERIFIED_CODE,
				Message: enums.ACCOUNT_NOT_VERIFIED_MESS,
			},
		}, nil
	}

	token, err := u.generateToken(utils.GenerateUniqueUUid(), userName, utils.GetCurrentTimestamp())
	if err != nil {
		return &model.ResponseLogin{
			Result: model.Result{
				Code:    enums.ACCOUNT_NOT_EXIST_CODE,
				Message: enums.ACCOUNT_NOT_EXIST_MESS,
			},
		}, nil
	}
	return &model.ResponseLogin{
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		JwtToken: token,
	}, nil
}

func (u *JwtUseCase) VerifiedAccount(ctx context.Context, userName string, code string) (*model.VerifiedAccountResp, error) {
	account, err := u.userRepositoryPort.GetInfomationByUserName(ctx, userName)
	if err != nil {
		return &model.VerifiedAccountResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}

	if account == nil {
		return &model.VerifiedAccountResp{
			Result: model.Result{
				Code:    enums.ACCOUNT_NOT_EXIST_CODE,
				Message: enums.ACCOUNT_NOT_EXIST_MESS,
			},
		}, nil
	}

	numOffset, err := strconv.Atoi(code)
	if err != nil {
		return &model.VerifiedAccountResp{
			Result: model.Result{
				Code:    enums.CONVERT_TO_NUMBER_CODE,
				Message: enums.CONVERT_TO_NUMBER_MESS,
			},
		}, nil
	}
	if account.OTPCode != int64(numOffset) {
		return &model.VerifiedAccountResp{
			Result: model.Result{
				Code:    enums.VERIFIED_ACCOUNT_ERROR_CODE,
				Message: enums.VERIFIED_ACCOUNT_ERROR_MESS,
			},
		}, nil
	}
	err = u.userRepositoryPort.UpdateAccount(ctx, &model.Account{
		ID:         account.ID,
		IsVerified: enums.IS_VERIFIED,
	})
	if err != nil {
		return &model.VerifiedAccountResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &model.VerifiedAccountResp{
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
func (u *JwtUseCase) ResendOtp(ctx context.Context, email, user_name string) (*model.ResendOtpResp, error) {

	newCode := utils.CodeOPT()

	account, err := u.userRepositoryPort.GetInfomationByUserName(ctx, user_name)

	if err != nil {
		return &model.ResendOtpResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	if account == nil {
		return &model.ResendOtpResp{
			Result: model.Result{
				Code:    enums.ACCOUNT_NOT_EXIST_CODE,
				Message: enums.ACCOUNT_NOT_EXIST_MESS,
			},
		}, nil
	}
	err = utils.SendEmail(email, int64(newCode))
	if err != nil {
		return &model.ResendOtpResp{
			Result: model.Result{
				Code:    enums.SEND_EMAIL_CODE_ERROR,
				Message: enums.SEND_EMAIL_MESS_ERROR,
			},
		}, nil
	}
	err = u.userRepositoryPort.UpdateAccount(ctx, &model.Account{
		ID:         account.ID,
		IsVerified: enums.IS_VERIFIED,
		Email:      email,
		OTPCode:    int64(newCode),
	})
	if err != nil {
		return &model.ResendOtpResp{
			Result: model.Result{
				Code:    enums.DB_ERR_CODE,
				Message: enums.DB_ERR_MESS,
			},
		}, nil
	}
	return &model.ResendOtpResp{
		Result: model.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
	}, nil
}
