package router

import (
	"learnGin/src/api/controllers"
	common "learnGin/src/common/middleware"
	socketIO "learnGin/src/libs/socket"

	"github.com/gin-gonic/gin"
)

func InitPublicRoute(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		api.POST("/auth/register", controllers.Registor)
		api.POST("/auth/login", controllers.Login)
		api.POST("/auth/mail-otp", controllers.VerifyMailOTP)
		api.POST("/auth/phone-otp", controllers.VerifyPhoneOTP)
	}
}

func InitPrivateRouter(r *gin.Engine) {

	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	api := r.Group("/api")
	{
		// User
		api.Use(common.Middleware())
		api.Use(common.Intercepter())
		api.GET("/users", controllers.GetUsers)

		api.POST("/file/upload", controllers.UploadFile)
	}
}

func InitSocketRoute(r *gin.Engine) {
	socketRouter := r.Group("/")
	socketRouter.GET("/socket.io/*any", gin.WrapH(socketIO.SocketServer))
	socketRouter.POST("/socket.io/*any", gin.WrapH(socketIO.SocketServer))
}
