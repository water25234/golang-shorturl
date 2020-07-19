package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/water25234/golang-shorturl/app/controller"
	core "github.com/water25234/golang-shorturl/config"
	"github.com/water25234/golang-shorturl/server"
)

var maxAttempts int
var decayMinutes int

type ThrottleDetail struct {
	Time      string
	IpAddress string
	UrlPath   string
}

func init() {
	maxAttempts = 10
	decayMinutes = 1
}

func ExecuteThrottle() gin.HandlerFunc {

	return func(c *gin.Context) {

		keysPattern := core.GetServerConfig().IpAddress + ":*"
		keysArray := server.GetKeys(keysPattern)
		keysLen := len(keysArray) + 1

		if keysLen > maxAttempts {
			ThrottleCount := gin.H{
				"ThrottleCount": keysLen,
			}
			//c.AbortWithStatusJSON(http.StatusTooManyRequests, api.GetErrorResponse(ThrottleCount, "Too many attempts, please slow down the request."))
			c.HTML(http.StatusOK, "index.tmpl", controller.GetErrorResponse(ThrottleCount, "Too many attempts, please slow down the request."))
			c.Abort()
			return
		}
		time := time.Now().UTC().Format("2006-01-02 03:04:05")
		key := core.GetServerConfig().IpAddress + ":" + time
		decaysecond := decayMinutes * 60

		ThrottleDetail := ThrottleDetail{
			Time:      time,
			IpAddress: core.GetServerConfig().IpAddress,
			UrlPath:   c.Request.URL.Path,
		}

		json, err := json.Marshal(ThrottleDetail)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, controller.GetErrorResponse(nil, "json fail."))
			return
		}
		value := string(json)

		server.SetRedis(key, value, decaysecond)
		c.Set("ThrottleCount", keysLen)
	}
}
