package main

import (
	"bcc-be-freepass-2024/database"
	"bcc-be-freepass-2024/handler"
	"bcc-be-freepass-2024/repository"
	"bcc-be-freepass-2024/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error when loading .env file: " + err.Error())
	}

	db, err := database.MakeConnection(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		log.Fatalln("Error when connecting to database: " + err.Error())
	}

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/register", userHandler.Register)
	v1.POST("/login", userHandler.Login)

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalln(err)
	}
}
