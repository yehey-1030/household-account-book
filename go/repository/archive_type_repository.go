package repository

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository/database"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
)

type ArchiveTypeRepository interface {
	List(ctx context.Context) ([]domain.ArchiveType, error)
	GetById(ctx context.Context, id int) (domain.ArchiveType, error)
}
type archiveTypeRepository struct {
	archiveTypeSearcher database.ArchiveTypeSearcher
}

func NewArchiveTypeRepository(archiveTypeSearcher database.ArchiveTypeSearcher) ArchiveTypeRepository {
	return &archiveTypeRepository{archiveTypeSearcher: archiveTypeSearcher}
}

func (a *archiveTypeRepository) List(ctx context.Context) ([]domain.ArchiveType, error) {
	list, err := a.archiveTypeSearcher.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return list, nil
}

func (a *archiveTypeRepository) GetById(ctx context.Context, id int) (domain.ArchiveType, error) {
	archiveType, err := a.archiveTypeSearcher.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return archiveType, nil
}
