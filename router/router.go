package router

import (
	"github.com/gin-gonic/gin"
	webv1shortener "github.com/water25234/golang-shorturl/app/controller/web/v1/shortener"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("app/view/*")

	shortener := router.Group("/shortener")
	{
		shortener.GET("", webv1shortener.PageStateShortener)
	}

	return router
}
