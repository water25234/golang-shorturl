package apiv1shortener

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/water25234/golang-shorturl/app/controller"
	servershortener "github.com/water25234/golang-shorturl/app/server/shortener"
)

type SaveShortenerForm struct {
	Url string `json:"url" form:"url"`
}

func SaveShortenerURL(ctx *gin.Context) {

	saveShortenerForm := &SaveShortenerForm{}

	ctx.Bind(&saveShortenerForm)

	if saveShortenerForm.Url == "" {
		ctx.JSON(http.StatusUnauthorized, controller.GetSuccessResponse("request parameter is failure"))
	}

	ch := make(chan interface{})
	go func() {
		ch <- servershortener.SaveShortenerURL(saveShortenerForm.Url)
	}()

	response := <-ch

	ctx.JSON(http.StatusOK, controller.GetSuccessResponse(response))
	ctx.Abort()
	return
}

func GetShortenerURL(ctx *gin.Context) {
	shortenerID := ctx.Param("shortenerID")

	if shortenerID == "" {
		ctx.JSON(http.StatusUnauthorized, controller.GetSuccessResponse("request parameter is failure"))
	}

	shortenerURL := servershortener.GetShortenerURL(shortenerID)

	ctx.JSON(http.StatusOK, controller.GetSuccessResponse(gin.H{"shortenerURL": shortenerURL}))
}
