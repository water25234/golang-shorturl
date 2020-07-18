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
		ctx.JSON(http.StatusUnauthorized, controller.GetSuccessResponse("request parameter is failure"))
	}

	response := servershortener.SaveShortenerURL(saveShortenerForm.Url)

	ctx.JSON(http.StatusOK, controller.GetSuccessResponse(response))
}

func GetShortenerURL(ctx *gin.Context) {
	shortenerID := ctx.Param("shortenerID")

	if shortenerID == "" {
		ctx.JSON(http.StatusUnauthorized, controller.GetSuccessResponse("request parameter is failure"))
	}

	shortenerURL := servershortener.GetShortenerURL(shortenerID)

	ctx.JSON(http.StatusOK, controller.GetSuccessResponse(gin.H{"shortenerURL": shortenerURL}))
}
