package verification

import (
	"mospol/database/postgres"
	"mospol/internal/entity"
	generator "mospol/internal/functions/jwt_generator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Verify(ctx *gin.Context, pg postgres.PostgresDB) (b bool) {
	cookie, err := ctx.Cookie("articles_service")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, entity.ErrorResponse{Error: "unauthorized"})
		return
	}

	user, password := generator.Parser(cookie)

	res, err := pg.CheckAuthroEmail(user, password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, entity.ErrorResponse{Error: "can`t load db"})
		return
	}

	if !res {
		ctx.JSON(http.StatusNotFound, entity.ErrorResponse{Error: "User not found"})
		return
	}

	b = true
	return
}
