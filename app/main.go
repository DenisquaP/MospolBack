package main

import (
	"io"
	"log"
	"mospol/database/postgres"
	_ "mospol/docs"
	"mospol/internal/services"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title	articles_api
// @version	1.0
// @description	A service to create, read and comment articles

// @host	localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/register", services.Register)
	router.POST("/auth", services.Auth)
	router.POST("/logout", services.LogOut)

	router.POST("/new_article", services.CreateAtricle)
	router.POST("/new_comment", services.CreateComment)

	router.GET("/get_articles", services.GetArticles)
	router.GET("/get_article", services.GetArticle)

	if err := pg.MigrationsUp(); err != nil {
		if err.Error() != "no change" {
			log.Fatal(err)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
