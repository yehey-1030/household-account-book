package setting

import (
	"github.com/gin-gonic/gin"
	"github.com/yehey-1030/household-account-book/go/handler"
	gormlib "gorm.io/gorm"
)

func GetMiddleWares(startupMessage, version string, db *gormlib.DB) ([]gin.HandlerFunc, []handler.Router) {
	handlerFuncs := []gin.HandlerFunc{}
	routers := []handler.Router{}
	return handlerFuncs, routers
}
