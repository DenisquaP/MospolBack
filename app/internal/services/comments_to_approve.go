package services

import (
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ToApprove godoc
// @Summary Returns all unapproved comments
// @Description Gets all comments
// @Proguce application/json
// @Success	200 {object} []postgres.UnapprovedComment
// @Router /all_comments [get]
func ToApprove(ctx *gin.Context) {
	pg, err := postgres.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Connection()
	if err != nil {
		log.Fatal(err)
	}

	defer pg.Close()

	comments, err := pg.GetComments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{Error: "can`t read comments"})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
