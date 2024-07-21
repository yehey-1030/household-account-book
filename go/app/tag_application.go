package app

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/app/request"
	"github.com/yehey-1030/household-account-book/go/app/response"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/service"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
)

type TagApplication interface {
	ListByParent(ctx context.Context, parentId int) (response.TagListResponse, error)
	ListRootTag(ctx context.Context, archiveTypeId int) (response.TagListResponse, error)
	Create(ctx context.Context, req request.CreateTagRequest) (response.TagResponse, error)
}
type tagApplication struct {
	tagService service.TagService
}

func NewTagApplication(tagService service.TagService) TagApplication {
	return &tagApplication{tagService: tagService}
}

func (t *tagApplication) ListByParent(ctx context.Context, parentId int) (response.TagListResponse, error) {
	list, err := t.tagService.ListByParent(ctx, parentId)
	if err != nil {
		return response.TagListResponse{}, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}

	var tags response.TagListResponse
	for _, tag := range list {
		tags.Tags = append(tags.Tags, tagResponseFrom(tag))
	}
	return tags, nil
}

func (t *tagApplication) ListRootTag(ctx context.Context, archiveTypeId int) (response.TagListResponse, error) {
	list, err := t.tagService.ListRootTag(ctx, archiveTypeId)
	if err != nil {
		return response.TagListResponse{}, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}

	var tags response.TagListResponse
	for _, tag := range list {
		tags.Tags = append(tags.Tags, tagResponseFrom(tag))
	}
	return tags, nil
}

func (t *tagApplication) Create(ctx context.Context, req request.CreateTagRequest) (response.TagResponse, error) {
	toCreate := domain.NewTag(0, req.Name, req.ParentId, req.ArchiveTypeId)

	created, err := t.tagService.Create(ctx, toCreate)
	if err != nil {
		return response.TagResponse{}, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return tagResponseFrom(created), nil
}

func tagResponseFrom(tag domain.Tag) response.TagResponse {
	return response.TagResponse{
		TagId:         tag.Id(),
		Name:          tag.Name(),
		ArchiveTypeId: tag.ArchiveTypeId(),
		ParentId:      tag.ParentId(),
	}
}
