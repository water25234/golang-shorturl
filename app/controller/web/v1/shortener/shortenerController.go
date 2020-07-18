package webv1shortener

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/water25234/golang-shorturl/app/controller"
	servershortener "github.com/water25234/golang-shorturl/app/server/shortener"
)

func PageStateShortener(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		controller.GetSuccessResponse(gin.H{
			"ThrottleCount": 1,
		}))
}

func GetShortenerURL(ctx *gin.Context) {
	shortenerID := ctx.Param("shortenerID")

	if shortenerID == "" {
		ctx.JSON(http.StatusUnauthorized, controller.GetSuccessResponse("request parameter is failure"))
	}

	shortenerURL := servershortener.GetShortenerURL(shortenerID)

	ctx.Redirect(http.StatusFound, shortenerURL)
	ctx.Abort()
	return
}
