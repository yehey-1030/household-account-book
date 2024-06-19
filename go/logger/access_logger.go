package logger

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

var GinLoggerConfig = gin.LoggerConfig{
	Formatter: logFormatter,
	Output:    os.Stdout,
	SkipPaths: nil,
}

var skipPaths = []string{"/ping", "/resources"}

var timeZoneKST = time.FixedZone("KST", 9*60*60)

type AccessLog struct {
	ServiceId     string `json:"service_id" example:"cloud-iam"`
	Message       string `json:"message" example:""`
	Host          string `json:"host" example:""`
	Ip            string `json:"ip" example:""`
	HttpSessionId string `json:"http_session_id" example:""`
	RequestTime   string `json:"req_time" example:""`
	Uri           string `json:"uri" example:""`
	StatusCode    int    `json:"status_code" example:""`
	Latency       string `json:"latency" example:""`
	UserAgent     string `json:"user_agent" example:""`
	Referer       string `json:"referer" example:""`
	UserId        string `json:"user_id" example:""`
}

func logFormatter(param gin.LogFormatterParams) string {
	for _, skipPath := range skipPaths {
		if strings.HasPrefix(param.Path, skipPath) {
			return ""
		}
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	clientIp := "-"
	if v, ok := param.Keys["clientIp"]; ok && v != nil {
		clientIp = v.(string)
	}

	userAgent := "-"
	if v, ok := param.Keys["Origin-User-Agent"]; ok && v != nil {
		userAgent = v.(string)
	}

	referer := "-"
	if v := param.Request.Referer(); v != "" {
		referer = v
	}

	accessLog := AccessLog{
		ServiceId:   "cloud-iam-access",
		Message:     "-",
		Host:        param.Request.Host,
		Ip:          clientIp,
		RequestTime: param.TimeStamp.UTC().Format("2006-01-02T15:04:05.000"),
		Uri:         fmt.Sprintf("%s %s", param.Method, param.Path),
		StatusCode:  param.StatusCode,
		Latency:     param.Latency.String(),
		UserAgent:   userAgent,
		Referer:     referer,
	}

	log, err := json.Marshal(accessLog)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s \n", log)
}
