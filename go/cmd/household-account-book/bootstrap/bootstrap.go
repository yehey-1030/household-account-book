package bootstrap

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
	"github.com/yehey-1030/household-account-book/go/cmd/household-account-book/setting"
	"github.com/yehey-1030/household-account-book/go/constants/config"
	"github.com/yehey-1030/household-account-book/go/logger"
	"time"
)

func Init() {
	logger.InitServerLogger()
	logrus.AddHook(logger.NewLogHook())

	envInit()
}

func envInit() {
	_ = env.Parse(&config.DBConfigInfo)
}

func Run(startupMessage, version string) {

	db := setting.SetUpDB()
	fmt.Printf("db setup finished at %s\n", time.Now().UTC().Format(time.RFC3339))

	_ = setting.GetMiddleWares(startupMessage, version, db)
}
