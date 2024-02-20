package services

import (
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	"mospol/internal/functions/verification"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Article godoc
// @Summary To create an article
// @Description Creates an entry in db
// @Param tags body entity.CreateAtricleRequest true "CreateArticle"
// @Proguce application/json
// @Success	201
// @Router /new_article [post]
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
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: "can`t parse body"})
		return
	}

	// check jwt in cookie
	verification.Verify(ctx, pg)

	if err := pg.WriteAtricle(request); err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{Error: "can`t create an entry in db"})
		return
	}

	ctx.JSON(http.StatusCreated, entity.OkResponse{Message: "Created"})
}
