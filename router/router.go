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
	}

	// api
	v1 := router.Group("/api/v1")
	{
		authRouting := v1.Group("/shortener")
		{
			authRouting.POST("/save", apiv1shortener.SaveShortenerURL)
		}
	}

	return router
}
