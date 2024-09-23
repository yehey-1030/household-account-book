package service

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
)

type StatisticService interface {
	Total(ctx context.Context, req domain.StatisticByTypeQuery) (int, error)
}

type statisticService struct {
	statisticRepository repository.StatisticRepository
}

func NewStatisticService(statisticRepository repository.StatisticRepository) StatisticService {
	return &statisticService{statisticRepository: statisticRepository}
}

func (s *statisticService) Total(ctx context.Context, req domain.StatisticByTypeQuery) (int, error) {
	total, err := s.statisticRepository.Total(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return total, nil
}