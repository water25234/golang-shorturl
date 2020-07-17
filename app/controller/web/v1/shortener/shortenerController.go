package webv1shortener

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/water25234/golang-shorturl/app/controller"
)

func PageStateShortener(ctx *gin.Context) {
	ctx.HTML(
		http.StatusOK,
		"index.html",
		controller.GetSuccessResponse(gin.H{
			"ThrottleCount": 1,
		}))
}
