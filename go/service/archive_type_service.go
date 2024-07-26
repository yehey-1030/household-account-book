package service

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
)

type ArchiveTypeService interface {
	List(ctx context.Context) ([]domain.ArchiveType, error)
	GetById(ctx context.Context, id int) (domain.ArchiveType, error)
}

type archiveTypeService struct {
	archiveTypeRepository repository.ArchiveTypeRepository
}

func NewArchiveTypeService(archiveTypeRepository repository.ArchiveTypeRepository) ArchiveTypeService {
	return &archiveTypeService{archiveTypeRepository: archiveTypeRepository}
}

func (a *archiveTypeService) List(ctx context.Context) ([]domain.ArchiveType, error) {
	list, err := a.archiveTypeRepository.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return list, nil
}

func (a *archiveTypeService) GetById(ctx context.Context, id int) (domain.ArchiveType, error) {
	archiveType, err := a.archiveTypeRepository.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return archiveType, nil
}
