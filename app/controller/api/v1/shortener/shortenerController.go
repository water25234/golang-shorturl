package apiv1shortener

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/water25234/golang-shorturl/app/controller"
	servershortener "github.com/water25234/golang-shorturl/app/server/shortener"
)

type SaveShortenerForm struct {
	Url string `form:url binding:"required"`
}

func SaveShortenerURL(ctx *gin.Context) {
	saveShortenerForm := &SaveShortenerForm{}

	ctx.BindJSON(&saveShortenerForm)

	if saveShortenerForm.Url == "" {
		ctx.JSON(http.StatusUnauthorized, controller.GetSuccessResponse("failure"))
	}

	response := servershortener.SaveShortenerURL(saveShortenerForm.Url)

	ctx.JSON(http.StatusOK, controller.GetSuccessResponse(response))
}
