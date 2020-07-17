package router

import (
	"github.com/gin-gonic/gin"
	webv1shortener "github.com/water25234/golang-shorturl/app/controller/web/v1/shortener"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("app/view/*")

	webv1 := router.Group("/web/v1")
	{
		authRouting := webv1.Group("/auth")
		{
			authRouting.GET("", webv1shortener.PageStateShortener)
		}
	}

	return router
}
