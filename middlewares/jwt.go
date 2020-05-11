package middlewares

import (
	"blog/pkg/e"
	"blog/pkg/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data interface{}
		code := e.Success
		// 获取tokent
		token := ctx.Query("token")
		if token == "" {
			token = GetTokenFromHeader(ctx)
		}
		if token == "" {
			code = e.InvalidParams
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.TokenExpired
				default:
					code = e.TokenInvalid
				}
			}
			// todo... check claims
		}

		if code != e.Success {
			ctx.JSON(401, gin.H{
				"code":    401,
				"message": e.GetMessage(code),
				"data":    data,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func GetTokenFromHeader(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		return ""
	}
	if strings.HasPrefix(bearerToken, "Bearer ") {
		return bearerToken[7 : len(bearerToken)-1]
	}
	return ""
}
