package servers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/yehey-1030/household-account-book/go/errmodel"
	"net/http"
	"strings"
)

func SendErrorWithErrorField(ctx *gin.Context, err error, addErrorField map[string]interface{}) {
	statusCode, errMsg, _ := errmodel.ErrResponseFrom(err)

	if len(addErrorField) > 0 {
		addErrorField["error"] = errMsg
		ctx.AbortWithStatusJSON(statusCode, addErrorField)
	} else {
		ctx.AbortWithStatusJSON(statusCode, gin.H{"error": errMsg})
	}

	_ = ctx.Error(err)

	if statusCode == http.StatusInternalServerError {
		logrus.WithContext(ctx).Errorf("%v", err)
	} else {
		logrus.WithContext(ctx).Warnf("%v", err)
	}
}

func SendError(ctx *gin.Context, err error) {
	SendErrorWithErrorField(ctx, err, nil)
}

func SendBindingError(ctx *gin.Context, bindingErr error) {
	var domainErr *errmodel.BadRequestError
	switch bindingErr.(type) {
	case validator.ValidationErrors:
		var messageList []string
		errs := bindingErr.(validator.ValidationErrors)
		for _, err := range errs {
			var message string
			message += "Validation failed on field '" + err.Field() + "'"
			message += ", condition: " + err.ActualTag()
			if err.Param() != "" {
				message += "=" + err.Param()
			}
			messageList = append(messageList, message)
		}

		domainErr = errmodel.NewBadRequestError(strings.Join(messageList, "\n"), nil)
	default:
		domainErr = errmodel.NewBadRequestError("", bindingErr)
	}
	SendError(ctx, domainErr)
}

func SendResponse(ctx *gin.Context, object interface{}, err error) {
	SendResponseWithErrorField(ctx, object, err, nil)
}

func SendResponseWithErrorField(ctx *gin.Context, object interface{}, err error, addErrorField map[string]interface{}) {
	method := strings.ToLower(ctx.Request.Method)
	var httpStatus int
	if method == "delete" {
		httpStatus = http.StatusNoContent
	} else {
		httpStatus = http.StatusOK
	}
	if err != nil {
		SendErrorWithErrorField(ctx, err, addErrorField)
	} else if object == nil {
		ctx.Status(httpStatus)

	} else {
		ctx.JSON(httpStatus, object)
	}
}

type DataResponse struct {
	Object     interface{} `json:"item,omitempty"`
	StatusCode int         `json:"status" example:"409"`
	Error      string      `json:"error,omitempty" example:"중복된 이름입니다."`
}

type MetaDataResponse struct {
	Total   int `json:"total" example:"10"`
	Success int `json:"success" example:"9"`
	Failure int `json:"failure" example:"1"`
}

type BulkResponse struct {
	DataResponse     []DataResponse   `json:"data"`
	MetaDataResponse MetaDataResponse `json:"metadata"`
}

func SendBulkResponse(ctx *gin.Context, appResults []errmodel.Result, err error) {
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}

	success := 0
	var data []DataResponse
	for _, r := range appResults {
		var obj interface{}
		if r.Object() != nil {
			obj = r.Object()
		}

		if r.ErrorObject() != nil {
			statusCode, errMsg, _ := errmodel.ErrResponseFrom(r.ErrorObject())
			data = append(data, DataResponse{obj, statusCode, errMsg})
			logrus.Warnf("bulk api %s. bulk err : %s", ctx.Request.RequestURI, r.ErrorObject())
		} else {
			success++
			data = append(data, DataResponse{obj, http.StatusOK, ""})
		}
	}

	metaData := MetaDataResponse{
		len(appResults),
		success,
		len(appResults) - success,
	}

	ctx.JSON(http.StatusOK, BulkResponse{data, metaData})
}

func SendIdentityResponse(ctx *gin.Context, resHeader map[string][]string, resBody string, statusCode int) {
	for k, v := range resHeader {
		if len(v) > 0 {
			ctx.Header(k, v[0])
		}
	}

	ctx.Status(statusCode)
	_, _ = ctx.Writer.Write([]byte(resBody))
}
