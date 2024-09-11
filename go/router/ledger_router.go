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

type LedgerRouter struct {
	ledgerApplication app.LedgerApplication
	prefix            string
}

func NewLegerRouter(ledgerApplication app.LedgerApplication) *LedgerRouter {
	return &LedgerRouter{ledgerApplication: ledgerApplication, prefix: "/api/v2"}
}

func (r *LedgerRouter) Routes() []handler.Route {
	return []handler.Route{
		handler.NewRoute(http.MethodGet, fmt.Sprintf("%s/ledgers", r.prefix), r.list),
		handler.NewRoute(http.MethodPost, fmt.Sprintf("%s/ledgers", r.prefix), r.create),
	}
}

func (r *LedgerRouter) list(ctx *gin.Context) {

	var req request.LedgerListRequest
	if err := ctx.ShouldBind(&req); err != nil {
		servers.SendBindingError(ctx, err)
		return
	}

	response, err := r.ledgerApplication.List(ctx, req)
	servers.SendResponse(ctx, response, err)
}

func (r *LedgerRouter) create(ctx *gin.Context) {
	var req request.LedgerCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		servers.SendBindingError(ctx, err)
		return
	}

	response, err := r.ledgerApplication.Create(ctx, req)
	servers.SendResponse(ctx, response, err)
}
