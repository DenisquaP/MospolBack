package services

import (
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	var request entity.CreateCommentRequest

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	if err := pg.Connection(); err != nil {
		log.Fatal("connection failed")
	}

	defer pg.Close()

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: "can`t parse body"})
		return
	}
}
