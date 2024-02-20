package services

import (
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApproveComment godoc
// @Summary To accept or delete comment
// @Description Updates an entry in db
// @Param tags body entity.ApproveRequest true "Approve"
// @Proguce application/json
// @Success	200 {object} entity.OkResponse
// @Router /approve_comment [patch]
func ApproveComment(ctx *gin.Context) {
	var request entity.ApproveRequest

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
		ctx.JSON(http.StatusOK, entity.ErrorResponse{Error: "can`t parse body"})
		return
	}

	err = pg.ApproveComment(request)
	if err != nil {
		ctx.JSON(http.StatusOK, entity.ErrorResponse{Error: "can`t update an entry in db"})
		return
	}

	ctx.JSON(http.StatusOK, entity.OkResponse{Message: "Updated"})
}
