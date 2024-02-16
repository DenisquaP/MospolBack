package services

import (
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArticle godoc
// @Summary		To get all articles from db
// @Description	Gets an entry from db by article_id
// @Proguce		application/json
// @Param		page query int true "Page number"
// @Success		200 {object} postgres.ArticleStr
// @Router		/get_article [get]
func GetArticle(ctx *gin.Context) {
	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	if err := pg.Connection(); err != nil {
		log.Fatal("connection failed")
	}

	defer pg.Close()

	article_id, err := strconv.Atoi(ctx.Query("article_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: "wrong article_id"})
		return
	}

	article, err := pg.ReadArticle(article_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: "can`t find in db"})
		return
	}

	ctx.JSON(http.StatusOK, article)
}
