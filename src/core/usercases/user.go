package usercases

import (
	"context"
	"ecommerce_site/src/adapter/model"
	"ecommerce_site/src/common/utils"
	"ecommerce_site/src/core/entities"
	"ecommerce_site/src/core/ports"
)

type UserUseCase struct {
	user ports.RepositoryUser
}

func NewUserUseCase(user ports.RepositoryUser) *UserUseCase {
	return &UserUseCase{
		user: user,
	}
}
func (a *UserUseCase) AddProfile(ctx context.Context, req *entities.UsersReq) (*entities.UsersResp, error) {
	id := utils.GenerateUniqueUUid()

	err := a.user.AddProfile(ctx, &model.Users{
		Id:      id,
		Name:    req.Name,
		Age:     req.Age,
		Address: req.Address,
	})
	if err != nil {
		return &entities.UsersResp{
			Result: entities.Result{
				Code:    1,
				Message: "err 1",
			},
		}, nil
	}
	return &entities.UsersResp{
		Id: id,
		Result: entities.Result{
			Code:    0,
			Message: "Sucess",
		},
	}, nil
}
