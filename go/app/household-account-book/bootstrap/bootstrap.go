package bootstrap

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yehey-1030/household-account-book/go/app/household-account-book/setting"
	"github.com/yehey-1030/household-account-book/go/logger"
	"time"
)

func Init() {
	logger.InitServerLogger()
	logrus.AddHook(logger.NewLogHook())

	envInit()
}

func envInit() {

}

func Run(startupMessage, version string) {

	db := setting.SetUpDB
	fmt.Println("db setup finished at %s", time.Now().UTC().Format(time.RFC3339))

	handlerFunc, routers := setting.GetMiddlewares(startupMessage, version, db)
}
