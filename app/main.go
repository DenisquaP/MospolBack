package main

import (
	"io"
	"log"
	"mospol/database/postgres"
	"mospol/internal/services"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	router.POST("/new_article", services.CreateAtricle)
	router.POST("/new_author", services.CreateAuthor)
	router.POST("/new_comment", services.CreateComment)

	if err := pg.MigrationsUp(); err != nil {
		if err.Error() != "no change" {
			log.Fatal(err)
		}
	}

	router.Run()
}
