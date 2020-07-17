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

func GetShortenerURL(ctx *gin.Context, w http.ResponseWriter, r *http.Request) {
	shortenerID := ctx.Param("shortenerID")

	shortenerID = servershortener.GetShortenerURL(shortenerID)

	http.Redirect(w, r, "http://127.0.0.1:8080/shortener/"+shortenerID, 302)
}
