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

	if err := pg.MigrationsUp(); err.Error() != "no change" {
		log.Fatalf("Migration create was failed: %v", err)
	}

	router.Run()
}
