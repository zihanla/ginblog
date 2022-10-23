package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			//AllowAllOrigins:  true,
			AllowOrigins:  []string{"*"}, // 等同于允许所有域名 #AllowAllOrigins:  true
			AllowMethods:  []string{"*"},
			AllowHeaders:  []string{"*", "Authorization"},
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			//AllowCredentials: true,
			MaxAge: 12 * time.Hour,
		},
	)
}
