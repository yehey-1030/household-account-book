package repository

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository/database"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
)

type LedgerRepository interface {
	List(ctx context.Context, pagingQuery domain.LedgerPagingQuery) ([]domain.Ledger, error)
}

type ledgerRepository struct {
	ledgerSearcher database.LedgerSearcher
}

func NewLedgerSearcher(ledgerSearcher database.LedgerSearcher) LedgerRepository {
	return &ledgerRepository{ledgerSearcher: ledgerSearcher}
}

func (l *ledgerRepository) List(ctx context.Context, pagingQuery domain.LedgerPagingQuery) ([]domain.Ledger, error) {
	ledgers, err := l.ledgerSearcher.List(ctx, pagingQuery)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return ledgers, nil
}
