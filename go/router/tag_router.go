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

type TagRouter struct {
	tagApplication app.TagApplication
	prefix         string
}

func NewTagRouter(tagApplication app.TagApplication) *TagRouter {
	return &TagRouter{tagApplication: tagApplication, prefix: "/api/v2"}
}

func (r *TagRouter) Routes() []handler.Route {
	return []handler.Route{
		handler.NewRoute(http.MethodPost, fmt.Sprintf("%s/tags", r.prefix), r.create),
		handler.NewRoute(http.MethodGet, fmt.Sprintf("%s/tags/:parent_id/childs", r.prefix), r.listByParent),
	}
}

func (r *TagRouter) create(ctx *gin.Context) {
	var req request.CreateTagRequest
	if err := ctx.ShouldBind(&req); err != nil {
		servers.SendBindingError(ctx, err)
		return
	}

	response, err := r.tagApplication.Create(ctx, req)
	servers.SendResponse(ctx, response, err)
}

func (r *TagRouter) listByParent(ctx *gin.Context) {
	var parentId request.UriParentId
	if err := ctx.ShouldBindUri(&parentId); err != nil {
		servers.SendBindingError(ctx, err)
		return
	}

	response, err := r.tagApplication.ListByParent(ctx, parentId.ParentId)
	servers.SendResponse(ctx, response, err)
}
