package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yehey-1030/household-account-book/go/app"
	"github.com/yehey-1030/household-account-book/go/handler"
	"github.com/yehey-1030/household-account-book/go/handler/servers"
	"net/http"
)

type ArchiveTypeRouter struct {
	archiveTypeApplication app.ArchiveTypeApplication
	prefix                 string
}

func NewArchiveTypeRouter(archiveTypeApplication app.ArchiveTypeApplication) *ArchiveTypeRouter {
	return &ArchiveTypeRouter{archiveTypeApplication: archiveTypeApplication, prefix: "/api/v2"}
}

func (r *ArchiveTypeRouter) Routes() []handler.Route {
	return []handler.Route{
		handler.NewRoute(http.MethodGet, fmt.Sprintf("%s/archive-types", r.prefix), r.list),
	}
}

func (r *ArchiveTypeRouter) list(ctx *gin.Context) {
	response, err := r.archiveTypeApplication.List(ctx)
	servers.SendResponse(ctx, response, err)
}
