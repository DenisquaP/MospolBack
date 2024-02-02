package services

import (
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAtricle(ctx *gin.Context) {
	var request entity.CreateAtricleRequest

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}

	defer pg.Close()

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: err})
	}

	if err := pg.WriteAtricle(request); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: err})
	}

	ctx.JSON(http.StatusCreated, "ok")
}
