package services

import (
	"fmt"
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticles(ctx *gin.Context) {
	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	if err := pg.Connection(); err != nil {
		log.Fatal("connection failed")
	}

	defer pg.Close()

	art, err := pg.ReadArticles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{Error: "can`t get articles"})
		fmt.Println(err)
		return
	}

	if art == nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: "there is no articles"})
		return
	}

	var response []entity.GetAtricleResponse
	for _, a := range art {
		r := entity.GetAtricleResponse{
			Title:   a.Title,
			Content: a.Content,
			Author:  a.Author,
		}
		response = append(response, r)
	}

	ctx.JSON(http.StatusOK, response)
}
