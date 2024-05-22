package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type LogHook struct {
	ctx *gin.Context
}

func NewLogHook() *LogHook {
	return &LogHook{}
}

func (hook *LogHook) Levels() []logrus.Level { return logrus.AllLevels }

func (hook *LogHook) Fire(entry *logrus.Entry) error {
	entry.Data["service_id"] = "HAB-error"
	ctx := entry.Context
	if ctx == nil {
		return nil
	}

	userAgent := "-"
	if v := ctx.Value("Origin-User-Agent"); v != nil {
		userAgent = v.(string)
	}
	entry.Data["user_agent"] = userAgent

	clientIp := "-"
	if v := ctx.Value("clientIp"); v != nil {
		clientIp = v.(string)
	}
	entry.Data["ip"] = clientIp

	userId := "-"
	if v := ctx.Value("Principal"); v != nil {
		userId = v.(string)
	}
	entry.Data["user_id"] = userId

	if ginCtx, ok := ctx.(*gin.Context); ok {
		entry.Data["host"] = ginCtx.Request.Host
		entry.Data["uri"] = ginCtx.Request.Method + " " + ginCtx.Request.RequestURI

		statusCode := ""
		if ginCtx.Writer.Status() > 0 {
			statusCode = strconv.Itoa(ginCtx.Writer.Status())
		}
		entry.Data["status_code"] = statusCode

		requestBody := ""
		if v, ok := ginCtx.Get("request_body"); ok {
			requestBody = v.(string)
		}
		entry.Data["request_body"] = requestBody
	}

	return nil
}
