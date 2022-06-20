package main

import (
	"database/sql"
	"diary/auth"
	"diary/handler"
	"diary/user"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db, err := sql.Open("sqlite3", "diary.db")
	// db, err := gorm.Open(sqlite.Open("diary.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	router.Use(cors.Default())
	api := router.Group("api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/check-email", userHandler.CheckEmailRegister)
	api.POST("/upload-avatar", userHandler.UploadAvatar)
	router.Run()


}


// Layering :
// handler
// service
// repository
// db


// input dari user
// handler -> mapping inputan dari user menjadi struct input
// service : mapping dari struct input  diubah ke struct User(db)
// repositoryy
// db