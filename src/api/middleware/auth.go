package middleware

import (
	"ecommerce_site/src/common/log"
	"ecommerce_site/src/core/usecases"

	"github.com/gin-gonic/gin"
)

type MiddleWare struct {
	jwtUseCase *usecases.JwtUseCase
}

func NewMiddleware(
	jwtUseCase *usecases.JwtUseCase,

) *MiddleWare {
	return &MiddleWare{
		jwtUseCase: jwtUseCase,
	}
}
func (m *MiddleWare) Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		data, err := m.jwtUseCase.Decrypt(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Set("username", data.UserName)
		log.Infof("user name", data.UserName)
		context.Next()
	}
}
