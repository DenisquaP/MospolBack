package services

import (
	"fmt"
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	emailsender "mospol/internal/functions/email_sender"
	"mospol/internal/functions/verification"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Comment godoc
// @Summary To create a comment to an article
// @Description Creates an entry in db
// @Param tags body entity.CreateCommentRequest true "CreateComment"
// @Proguce application/json
// @Success	201
// @Router /new_comment [post]
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

	verification.Verify(ctx, pg)

	err = emailsender.Sender("piskarev.py@yandex.ru", "new comment")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: "can`t send an email"})
		return
	}

	if err := pg.WriteComment(request); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, entity.ErrorResponse{Error: "can`t create an entry in db"})
		return
	}

	ctx.JSON(http.StatusCreated, "ok")
}
