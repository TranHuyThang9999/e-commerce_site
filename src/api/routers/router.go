package routers

import (
	"ecommerce_site/src/api/controllers"
	"ecommerce_site/src/api/middleware"
	"ecommerce_site/src/configs"
	"fmt"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.elastic.co/apm/module/apmgin/v2"
	"go.elastic.co/apm/v2"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	controllersAccount *controllers.AccountController,
	controllerAuth *controllers.AuthController,
	middleware *middleware.MiddleWare,
	controllerProduct *controllers.ControllerProduct,
	cf *configs.Configs,
) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()
	tracer, err := apm.NewTracer("cms-backend", "v0.0.1")
	if err != nil {
		fmt.Errorf("error", err)
	}
	engine.Use(gin.Logger())
	engine.Use(apmgin.Middleware(engine, apmgin.WithTracer(tracer)))
	engine.Use(cors.AllowAll())

	r := engine.RouterGroup.Group("/sell")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/add", controllersAccount.CreateAccount)
	r.POST("/login", controllerAuth.Login)
	r.POST("/verified", controllerAuth.VerifiedAccount)
	r.POST("/resendOtp", controllerAuth.ResendOtp)
	r.POST("/product/add", controllerProduct.AddProduct)
	// userGroup.Use(middleware.Authenticate())
	// {

	// }

	return &ApiRouter{
		Engine: engine,
	}
}
