package router

import (
	"github.com/gin-gonic/gin"
	apiv1shortener "github.com/water25234/golang-shorturl/app/controller/api/v1/shortener"
	webv1shortener "github.com/water25234/golang-shorturl/app/controller/web/v1/shortener"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("app/view/*")

	// web
	shortener := router.Group("/shortener")
	{
		shortener.GET("", webv1shortener.PageStateShortener)
		shortener.GET("/:shortenerID", webv1shortener.GetShortenerURL)
	}

	// api
	v1 := router.Group("/api/v1")
	{
		shortenerRouting := v1.Group("/shortener")
		{
			shortenerRouting.GET("/:shortenerID", apiv1shortener.GetShortenerURL)
			shortenerRouting.POST("", apiv1shortener.SaveShortenerURL)
		}
	}

	return router
}
