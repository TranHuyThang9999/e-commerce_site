package routers

import (
	"ecommerce_site/src/adapter/configs"
	"ecommerce_site/src/api/controllers"
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
	controllersUser *controllers.ControllersUser,
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

	r.POST("/add", controllersUser.AddProfile)

	return &ApiRouter{
		Engine: engine,
	}
}
