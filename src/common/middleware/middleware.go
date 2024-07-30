package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"learnGin/src/api/services"
	common "learnGin/src/common/customError"
	customJWT "learnGin/src/libs/jwt"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		header := ctx.Request.Header

		token := header.Get("Authorization")
		token = token[7:]
		fmt.Println("token", token)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, common.ResponseError(common.UNAUTHORIZE, "Unauthorized"))
			return
		}

		userVerified := customJWT.VerifyJwt(ctx, token)
		// // Verify if user exist
		user := services.RunAuth(ctx, userVerified.Id)
		fmt.Println("========= verified user =====", user.Id.Hex())
		ctx.Header("userId", user.Id.Hex())
		ctx.Header("userName", user.Name)
		ctx.Header("Email", user.Email)
		ctx.Next()
	}
}
