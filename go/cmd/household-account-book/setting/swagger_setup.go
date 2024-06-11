package setting

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func InitSwagger(r *gin.Engine, endpointHost string) {
	docs.SwaggerInfo.Schemes = []string{"https"}
	docs.SwaggerInfo.Host = endpointHost

	//aclInternalApiList := map[string]map[string]bool{
	//	"/swagger/**": constants.AllMethod,
	//}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
