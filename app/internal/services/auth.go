package services

import (
	"log"
	"mospol/database/postgres"
	"mospol/internal/entity"
	generator "mospol/internal/functions/jwt_generator"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth godoc
// @Summary To auth user
// @Description Creates an entry in cookie
// @Param tags body entity.AuthRequest true "Auth"
// @Proguce application/json
// @Success	200 {object} entity.AuthResponse
// @Router /auth [post]
func Auth(ctx *gin.Context) {
	var request entity.AuthRequest

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

	cookie, err := generator.JwtGenerator(request.Email, request.Password)
	if err != nil {
		log.Fatal(err)
	}

	author, err := pg.GetAuthor(request.Email)
	if err != nil {
		log.Fatal(err)
	}

	ctx.SetCookie("articles_service", cookie, 3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusCreated, entity.AuthResponse{User: author.User, IsModerator: author.IsModerator})
}
