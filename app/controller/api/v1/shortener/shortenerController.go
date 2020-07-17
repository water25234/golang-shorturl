package apiv1shortener

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/water25234/golang-shorturl/app/controller"
)

type SaveShortenerForm struct {
	Url string `form:url binding:"required"`
}

func SaveShortenerUrl(ctx *gin.Context) {
	saveShortenerForm := &SaveShortenerForm{}
	if ctx.BindJSON(&saveShortenerForm) == nil {
		ctx.JSON(http.StatusUnauthorized, controller.GetSuccessResponse("failure"))
	}

	fmt.Println(saveShortenerForm.Url)

	if saveShortenerForm.Url == "" {
		ctx.JSON(http.StatusUnauthorized, controller.GetSuccessResponse("failure"))
	}

	ctx.JSON(http.StatusOK, controller.GetSuccessResponse(gin.H{
		"ThrottleCount": 1,
		"userId":        1,
	}))
}
