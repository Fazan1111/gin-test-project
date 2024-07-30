package main

import (
	"fmt"
	envconfig "learnGin/src/common/envConfig"
	redisLib "learnGin/src/libs/redis"
	socketIO "learnGin/src/libs/socket"
	"learnGin/src/loader/mongo"
	"learnGin/src/router"
	"log"
	"net/http"
	"time"

	_ "learnGin/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	if err := mongo.ConnectToMongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	port := envconfig.GetEnv("APP_PORT")
	appMode := envconfig.GetEnv("APP_MODE")
	fmt.Println("env port ", port)
	if port == "" {
		port = "9000"
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	if appMode == "development" {
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// Initialize the Socket.IO server

	socketIO.InitSocket()
	socketIO.SocketEvents()

	// Initialize gin router
	router.InitPublicRoute(r)
	router.InitPrivateRouter(r)

	go socketIO.SocketServer.Serve()
	defer socketIO.SocketServer.Close()
	router.InitSocketRoute(r)

	// Connect redis
	redisLib.InitRedis()

	// Custom HTTP configuration
	s := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}
