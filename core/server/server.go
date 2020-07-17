package coreserver

import (
	"github.com/gin-gonic/gin"
	"github.com/water25234/golang-shorturl/config"
	"github.com/water25234/golang-shorturl/router"
	"github.com/water25234/golang-shorturl/server"
)

var Router *gin.Engine

func init() {
	config.SetServerGonfig()
	config.SetAppConfig()
	server.InitRedis()
	server.InitDataBase()
}

func StartServer() {
	Router = router.SetupRouter()
	Router.Run()
}
