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

type StatisticApplication interface {
	TotalByArchiveType(ctx context.Context, archiveTypeId int, req request.StatisticDateRangeRequest) (response.TotalByArchiveType, error)
}

type statisticApplication struct {
	statisticService service.StatisticService
}

func NewStatisticApplication(statisticService service.StatisticService) StatisticApplication {
	return &statisticApplication{statisticService: statisticService}
}

func (s *statisticApplication) TotalByArchiveType(ctx context.Context, archiveTypeId int, req request.StatisticDateRangeRequest) (response.TotalByArchiveType, error) {
	reqQuery := domain.StatisticByTypeQuery{
		ArchiveTypeId: archiveTypeId,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
	}

	total, err := s.statisticService.Total(ctx, reqQuery)
	if err != nil {
		return response.TotalByArchiveType{}, fmt.Errorf("[%s] %W", ioutil.FuncName(), err)
	}

	return response.TotalByArchiveType{Total: total}, nil
}
