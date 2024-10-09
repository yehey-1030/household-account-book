package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yehey-1030/household-account-book/go/app"
	"github.com/yehey-1030/household-account-book/go/app/request"
	"github.com/yehey-1030/household-account-book/go/handler"
	"github.com/yehey-1030/household-account-book/go/handler/servers"
	"net/http"
)

type StatisticRouter struct {
	statisticApplication app.StatisticApplication
	prefix               string
}

func NewStatisticRouter(application app.StatisticApplication) *StatisticRouter {
	return &StatisticRouter{statisticApplication: application}
}

func (r *StatisticRouter) Routes() []handler.Route {
	return []handler.Route{
		handler.NewRoute(http.MethodGet, fmt.Sprintf("%s/types/:archivetype_id/total", r.prefix), r.totalByType),
	}
}

func (r *StatisticRouter) totalByType(ctx *gin.Context) {
	var typeId int
	if err := ctx.ShouldBindUri(&typeId); err != nil {
		servers.SendBindingError(ctx, err)
		return
	}

	var dateRangeRequest request.StatisticDateRangeRequest
	if err := ctx.ShouldBind(&dateRangeRequest); err != nil {
		servers.SendBindingError(ctx, err)
		return
	}

	response, err := r.statisticApplication.TotalByArchiveType(ctx, typeId, dateRangeRequest)
	servers.SendResponse(ctx, response, err)
}
