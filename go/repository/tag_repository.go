package repository

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository/database"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
)

type TagRepository interface {
	ListByParent(ctx context.Context, parentId int) ([]domain.Tag, error)
	ListRootTag(ctx context.Context, archiveTypeId int) ([]domain.Tag, error)
	Create(ctx context.Context, tag domain.Tag) (domain.Tag, error)
}

type tagRepository struct {
	tagSearcher database.TagSearcher
}

func NewTagRepository(tagSearcher database.TagSearcher) TagRepository {
	return &tagRepository{tagSearcher: tagSearcher}
}

func (t *tagRepository) ListByParent(ctx context.Context, parentId int) ([]domain.Tag, error) {
	list, err := t.tagSearcher.ListByParent(ctx, parentId)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return list, nil
}

func (t *tagRepository) ListRootTag(ctx context.Context, archiveTypeId int) ([]domain.Tag, error) {
	list, err := t.tagSearcher.ListRootTag(ctx, archiveTypeId)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return list, nil
}

func (t *tagRepository) Create(ctx context.Context, tag domain.Tag) (domain.Tag, error) {
	created, err := t.tagSearcher.Create(ctx, tag)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return created, nil
}
