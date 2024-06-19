package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"time"
)

func ReadTimeoutHandler(timeout time.Duration) gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Body != nil {
			context.Request.Body = ioutil.NewTimeoutReader(context.Request.Body, timeout)
		}
	}
}
