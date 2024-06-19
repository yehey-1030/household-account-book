package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/yehey-1030/household-account-book/go/constants/config"
	"net/http"
)

func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		isAllowedOrigin := false
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			if config.ServerConfigInfo.UrlAccessCheckEnabled {
				for _, url := range config.ServerConfigInfo.AccessGrantedUrls {
					if origin == url {
						isAllowedOrigin = true
						break
					}
				}
			} else {
				isAllowedOrigin = true
			}

			if isAllowedOrigin {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Methods", c.GetHeader("Access-Control-Request-Method"))
				c.Header("Access-Control-Allow-Headers", c.GetHeader("Access-Control-Request-Headers"))
				c.Header("Access-Control-Max-Age", "86400")
				c.Header("Access-Control-Allow-Credentials", "true")
				c.Header("Vary", "Origin")
			}

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
		}
		c.Next()
	}
}
