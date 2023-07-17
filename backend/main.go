package main

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "golang_task/docs"
	"log"
	"os"

	"golang_task/controller"
	"golang_task/database"
	"golang_task/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

// test function
func sayHello(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Hello world!")
}

func main() {
	loadEnv()
	loadDatabase()
	startServer()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.StoredImage{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Run @title Golang Init
// @version 1.0
// @description This is init
// @host 0.0.0.0:8000
// @BasePath /
// @schemes http
func startServer() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	router.GET("/hello", sayHello)
	router.GET("/get_last_images", controller.GetLastImages)
	router.POST("/negative_image", controller.NegativeImage)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("0.0.0.0" + ":" + os.Getenv("BACKEND_PORT"))
}
