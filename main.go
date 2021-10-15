package main

import (
	"datatables_go/handler"
	"datatables_go/users"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/serverside_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService, userRepository)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("api/v1")
	api.POST("/users", userHandler.GetAll)
	router.Run()
}
