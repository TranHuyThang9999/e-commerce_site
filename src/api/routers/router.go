package routers

import (
	"ecommerce_site/src/api/controllers"
	"ecommerce_site/src/api/middleware"
	"ecommerce_site/src/configs"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	controllersAccount *controllers.AccountController,
	controllerAuth *controllers.AuthController,
	middleware *middleware.MiddleWare,
	controllerProduct *controllers.ControllerProduct,
	controllerImage *controllers.FileController,
	cf *configs.Configs,
) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()

	engine.Use(gin.Logger())
	engine.Use(cors.AllowAll())

	r := engine.RouterGroup.Group("/sell")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/add", controllersAccount.CreateAccount)
	r.POST("/login", controllerAuth.Login)
	r.POST("/verified", controllerAuth.VerifiedAccount)
	r.POST("/resendOtp", controllerAuth.ResendOtp)

	userGroup := r.Group("/user")
	userGroup.Use(middleware.Authenticate())
	{
		userGroup.POST("/product/add", controllerProduct.AddProduct)
		userGroup.GET("/product/list", controllerProduct.GetListProduct)
		userGroup.DELETE("/image/describe/:id", controllerImage.DeleteImageById)
		userGroup.DELETE("/product/", controllerProduct.DeleteProductById)
	}

	return &ApiRouter{
		Engine: engine,
	}
}
