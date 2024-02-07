package services

import (
	"mospol/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogOut godoc
// @Summary To logout
// @Description Deletes an cookie entry
// @Proguce application/json
// @Success	200
// @Router /logout [post]
func LogOut(ctx *gin.Context) {
	ctx.SetCookie("articles_service", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, entity.OkResponse{Message: "logout"})
}
