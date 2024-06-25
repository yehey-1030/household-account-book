package service

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
)

type LedgerService interface {
	List(ctx context.Context, query domain.LedgerPagingQuery) ([]domain.Ledger, error)
}

type ledgerService struct {
	ledgerRepository repository.LedgerRepository
}

func NewLedgerService(ledgerRepository repository.LedgerRepository) LedgerService {
	return &ledgerService{ledgerRepository: ledgerRepository}
}

func (l *ledgerService) List(ctx context.Context, query domain.LedgerPagingQuery) ([]domain.Ledger, error) {
	ledgers, err := l.ledgerRepository.List(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return ledgers, nil
}
