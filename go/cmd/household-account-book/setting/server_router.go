package setting

import (
	"github.com/gin-gonic/gin"
	"github.com/yehey-1030/household-account-book/go/app"
	"github.com/yehey-1030/household-account-book/go/handler"
	"github.com/yehey-1030/household-account-book/go/handler/servers"
	"github.com/yehey-1030/household-account-book/go/logger"
	"github.com/yehey-1030/household-account-book/go/repository"
	"github.com/yehey-1030/household-account-book/go/repository/database"
	"github.com/yehey-1030/household-account-book/go/router"
	"github.com/yehey-1030/household-account-book/go/service"
	gormlib "gorm.io/gorm"
	"time"
)

func GetMiddleWares(startupMessage, version string, db *gormlib.DB) ([]gin.HandlerFunc, []handler.Router) {
	var ledgerSearcher = database.NewLedgerSearcher(db)

	var ledgerRepository = repository.NewLedgerSearcher(ledgerSearcher)

	var ledgerService = service.NewLedgerService(ledgerRepository)

	var ledgerApplication = app.NewLedgerApplication(ledgerService)

	handlerFuncs := []gin.HandlerFunc{
		gin.LoggerWithConfig(logger.GinLoggerConfig),
		servers.ReadTimeoutHandler(time.Second * 30),
		servers.CorsHandler(),
	}
	routers := []handler.Router{
		router.NewLegerRouter(ledgerApplication),
	}
	return handlerFuncs, routers
}
