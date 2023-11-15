package bootstrap

import (
	"ecommerce_site/src/adapter"
	"ecommerce_site/src/adapter/postgresql"
	"ecommerce_site/src/api/controllers"
	"ecommerce_site/src/api/middleware"
	"ecommerce_site/src/api/routers"
	"ecommerce_site/src/core/usecases"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadValidator()...),
		fx.Options(loadEngine()...),
	}
}
func loadUseCase() []fx.Option {
	return []fx.Option{
		//	fx.Provide(usercases.NewUserUseCase),
		fx.Provide(usecases.NewUseCaseAccount),
		fx.Provide(usecases.NewJwtUseCase),
	}
}

func loadValidator() []fx.Option {
	return []fx.Option{
		fx.Provide(validator.New),
	}
}
func loadEngine() []fx.Option {
	return []fx.Option{
		//	fx.Provide(controllers.NewControllerUser),
		fx.Provide(routers.NewApiRouter),
		fx.Provide(middleware.NewMiddleware),
		fx.Provide(controllers.NewBaseController),
		fx.Provide(controllers.NewControllerAccount),
		fx.Provide(controllers.NewAuthController),
	}
}
func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(adapter.NewpostgreDb),
		fx.Provide(postgresql.NewTransaction),
		fx.Provide(postgresql.NewAccountRepository),
		fx.Provide(postgresql.NewRoleRepository),
	}
}