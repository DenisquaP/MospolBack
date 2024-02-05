package services

import (
	"fmt"
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	"net/http"
	"strconv"

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

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: "page should be integer"})
		return
	}

	art, err := pg.ReadArticles(page)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{Error: "can`t get articles"})
		fmt.Println(err)
		return
	}

	if art == nil {
		ctx.JSON(http.StatusOK, entity.OkResponse{Message: "there is no articles"})
		return
	}

	var response entity.GetArticlesResponse
	for _, a := range art {
		r := entity.GetAtricleResponse{
			Title:   a.Title,
			Content: a.Content,
			Author:  a.Author,
		}
		response.Articles = append(response.Articles, r)
	}

	lastPage, err := pg.LastPage()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{Error: "can`t load pages"})
	}

	response.LastPage = lastPage

	ctx.JSON(http.StatusOK, response)
}
