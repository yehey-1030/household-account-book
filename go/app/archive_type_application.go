package app

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/app/response"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/service"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
)

type ArchiveTypeApplication interface {
	List(ctx context.Context) (response.ArchiveTypeListResponse, error)
}

type archiveTypeApplication struct {
	archiveTypeService service.ArchiveTypeService
}

func NewArchiveTypeApplication(archiveTypeService service.ArchiveTypeService) ArchiveTypeApplication {
	return &archiveTypeApplication{archiveTypeService: archiveTypeService}
}

func (a *archiveTypeApplication) List(ctx context.Context) (response.ArchiveTypeListResponse, error) {
	list, err := a.archiveTypeService.List(ctx)
	if err != nil {
		return response.ArchiveTypeListResponse{}, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}

	var archiveTypes response.ArchiveTypeListResponse
	for _, archiveType := range list {
		archiveTypes.ArchiveTypes = append(archiveTypes.ArchiveTypes, archiveTypeResponseFrom(archiveType))
	}
	return archiveTypes, nil
}

func archiveTypeResponseFrom(archiveType domain.ArchiveType) response.ArchiveTypeResponse {
	return response.ArchiveTypeResponse{
		Id:   archiveType.Id(),
		Name: archiveType.Name(),
	}
}
