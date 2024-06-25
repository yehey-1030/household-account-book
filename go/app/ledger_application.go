package app

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/app/request"
	"github.com/yehey-1030/household-account-book/go/app/response"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/service"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"time"
)

type LedgerApplication interface {
	List(ctx context.Context, query request.LedgerListRequest) (response.LedgerListResponse, error)
}

type ledgerApplication struct {
	ledgerService service.LedgerService
}

func NewLedgerApplication(ledgerService service.LedgerService) LedgerApplication {
	return &ledgerApplication{ledgerService: ledgerService}
}

func (l *ledgerApplication) List(ctx context.Context, query request.LedgerListRequest) (response.LedgerListResponse, error) {
	pagingQuery := domain.LedgerPagingQuery{
		StartDate:     query.StartDate,
		EndDate:       query.EndDate,
		ArchiveTypeId: query.ArchiveTypeId,
		TagId:         0,
	}

	list, err := l.ledgerService.List(ctx, pagingQuery)
	if err != nil {
		return response.LedgerListResponse{}, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}

	var ledgers response.LedgerListResponse
	for _, ledger := range list {
		ledgers.LedgerList = append(ledgers.LedgerList, ledgerResponseFrom(ledger))
	}
	return ledgers, nil
}

func ledgerResponseFrom(ledger domain.Ledger) response.LedgerResponse {
	date, _ := time.Parse(time.DateOnly, ledger.Date().Format(time.DateOnly))
	return response.LedgerResponse{
		Title:  ledger.Title(),
		Memo:   ledger.Memo(),
		Amount: ledger.Amount(),
		Date:   date,
	}
}
