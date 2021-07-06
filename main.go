package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kholid74/golang/jwt-go/controllers"
	"github.com/kholid74/golang/jwt-go/database"
	"github.com/kholid74/golang/jwt-go/middleware"
	"github.com/kholid74/golang/jwt-go/models"
)

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
	r.Run(":8080")
}
