package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kholid74/desaku-auth/controllers"
	"github.com/kholid74/desaku-auth/database"
	"github.com/kholid74/desaku-auth/middleware"
	"github.com/kholid74/desaku-auth/models"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
		}

		protected := api.Group("/protected").Use(middleware.Authz())
		{
			protected.GET("/profile", controllers.Profile)
			protected.GET("/penduduk", controllers.Penduduk)
		}
	}

	return r
}

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	database.GlobalDB.AutoMigrate(&models.User{}, &models.Penduduk{})

	r := setupRouter()
	r.Run(":" + viper.GetString("APP_PORT"))
}
