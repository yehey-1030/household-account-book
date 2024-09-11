package app

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/app/request"
	"github.com/yehey-1030/household-account-book/go/app/response"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/service"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"github.com/yehey-1030/household-account-book/go/util/timeutil"
	"time"
)

type LedgerApplication interface {
	List(ctx context.Context, query request.LedgerListRequest) (response.LedgerListResponse, error)
	Create(ctx context.Context, req request.LedgerCreateRequest) (response.LedgerResponse, error)
}

type ledgerApplication struct {
	ledgerService service.LedgerService
}

func NewLedgerApplication(ledgerService service.LedgerService) LedgerApplication {
	return &ledgerApplication{ledgerService: ledgerService}
}

func (l *ledgerApplication) List(ctx context.Context, query request.LedgerListRequest) (response.LedgerListResponse, error) {
	pagingQuery := domain.LedgerPagingQuery{
		StartDate: query.StartDate,
		EndDate:   query.EndDate,
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

func (l *ledgerApplication) Create(ctx context.Context, req request.LedgerCreateRequest) (response.LedgerResponse, error) {
	targetDate := timeutil.StringToDate(req.Date)
	archiveType := domain.NewArchiveType(req.ArchiveTypeId, "")
	var tags []domain.Tag
	for _, t := range req.Tags {
		tags = append(tags, domain.NewTag(t, "", 0, 0))
	}

	toCreate := domain.NewLedger(0, req.Amount, req.Title, req.Memo, targetDate, req.IsExcluded, archiveType, tags)

	created, err := l.ledgerService.Create(ctx, toCreate)
	if err != nil {
		return response.LedgerResponse{}, fmt.Errorf("[%s] %w", ioutil.FuncName(), err)
	}
	return ledgerResponseFrom(created), nil
}

func ledgerResponseFrom(ledger domain.Ledger) response.LedgerResponse {
	date, _ := time.Parse(time.DateOnly, ledger.Date().Format(time.DateOnly))
	var tags []response.TagResponse
	for _, t := range ledger.Tags() {
		tags = append(tags, response.TagResponse{TagId: t.Id(), ArchiveTypeId: t.ArchiveTypeId(), ParentId: t.ParentId(), Name: t.Name()})
	}
	archiveType := response.ArchiveTypeResponse{Id: ledger.ArchiveType().Id(), Name: ledger.ArchiveType().Name()}
	return response.LedgerResponse{
		LedgerId:      ledger.Id(),
		Title:         ledger.Title(),
		Memo:          ledger.Memo(),
		Amount:        ledger.Amount(),
		Date:          date,
		ArchiveTypeId: archiveType,
		Tags:          tags,
	}
}
