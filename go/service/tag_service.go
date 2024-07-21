package service

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
)

type TagService interface {
	ListByParent(ctx context.Context, parentId int) ([]domain.Tag, error)
	ListRootTag(ctx context.Context, archiveTypeId int) ([]domain.Tag, error)
	Create(ctx context.Context, tag domain.Tag) (domain.Tag, error)
}

type tagService struct {
	tagRepository repository.TagRepository
}

func NewTagService(tagRepository repository.TagRepository) TagService {
	return &tagService{tagRepository: tagRepository}
}

func (t *tagService) ListByParent(ctx context.Context, parentId int) ([]domain.Tag, error) {
	list, err := t.tagRepository.ListByParent(ctx, parentId)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return list, nil
}

func (t *tagService) ListRootTag(ctx context.Context, archiveTypeId int) ([]domain.Tag, error) {
	list, err := t.tagRepository.ListRootTag(ctx, archiveTypeId)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return list, nil
}

func (t *tagService) Create(ctx context.Context, tag domain.Tag) (domain.Tag, error) {
	created, err := t.tagRepository.Create(ctx, tag)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return created, nil
}
