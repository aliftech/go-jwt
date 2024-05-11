package main

import (
	"github.com/aliftech/go-jwt/controllers"
	"github.com/aliftech/go-jwt/middleware"
	"github.com/aliftech/go-jwt/migrations"
	"github.com/aliftech/go-jwt/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.Setup()
	utils.ConnectDB()
	migrations.UserMigrate()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.ValidateAuth, controllers.Validate)
	r.Run()
}
