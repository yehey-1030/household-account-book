package bootstrap

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yehey-1030/household-account-book/go/cmd/household-account-book/setting"
	"github.com/yehey-1030/household-account-book/go/constants/config"
	"github.com/yehey-1030/household-account-book/go/handler"
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
	_ = env.Parse(&config.ServerConfigInfo)
}

func Run(startupMessage, version string) {

	db := setting.SetUpDB()
	fmt.Printf("db setup finished at %s\n", time.Now().UTC().Format(time.RFC3339))

	handlerFuncList, routers := setting.GetMiddleWares(startupMessage, version, db)

	var ginEngine = gin.New()
	setting.InitSwagger(ginEngine, "127.0.0.1:8000")

	ginEngine.Use(handlerFuncList...)
	_ = ginEngine.SetTrustedProxies(nil)

	rootRouter := handler.NewRootRouter(ginEngine, routers...)
	address := ""
	port := "8000"

	startupFinishedMessage := fmt.Sprintf("household-account-book-server start server at %s...", time.Now().UTC().Format(time.RFC3339))
	fmt.Println(startupFinishedMessage)

	handler.NewServer(rootRouter, address+":"+port).Start()
}
