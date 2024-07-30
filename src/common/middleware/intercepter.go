package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Intercepter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Pre request
		fmt.Println("Pre request")

		c.Next()

		fmt.Println("Post request")
	}
}
