package services

import (
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(ctx *gin.Context) {
	var request entity.CreateAuthorRequest

	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	if err := pg.Connection(); err != nil {
		log.Fatal(err)
	}

	defer pg.Close()

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: "can`t parse body"})
		return
	}

	if err := pg.WriteAuthor(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{Error: "can`t create an entry in db"})
		return
	}

	ctx.JSON(http.StatusCreated, "ok")
}
