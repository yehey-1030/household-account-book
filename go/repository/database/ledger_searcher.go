package database

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/lib/transaction/gorm_tx"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"gorm.io/gorm"
	"time"
)

type Ledger struct {
	LedgerId   int
	Amount     int
	Date       *time.Time
	Title      string
	Memo       string
	IsExcluded bool
}

func (l Ledger) TableName() string {
	return "ledger"
}

type LedgerSearcher interface {
	List(ctx context.Context, pagingQuery domain.LedgerPagingQuery) ([]domain.Ledger, error)
}

type ledgerSearcher struct {
	db *gorm.DB
}

func NewLedgerSearcher(db *gorm.DB) LedgerSearcher {
	return &ledgerSearcher{db: db}
}

func (l *ledgerSearcher) List(ctx context.Context, pagingQuery domain.LedgerPagingQuery) ([]domain.Ledger, error) {
	db := gorm_tx.FromContextWithDefault(ctx, l.db)

	result := db.Model(&Ledger{})
	//Where("ledger.archivetype_id = ?", pagingQuery.ArchiveTypeId)

	if !pagingQuery.StartDate.IsZero() {
		result = result.Where("ledger.date >= ?", pagingQuery.StartDate)
	}
	if !pagingQuery.EndDate.IsZero() {
		result = result.Where("ledger.date <= ?", pagingQuery.EndDate)
	}

	result = result.Order("ledger.date ASC")

	var ledgers []Ledger
	result = result.Find(&ledgers)
	if result.Error != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), result.Error)
	}

	var ledgerList []domain.Ledger
	for _, l := range ledgers {
		ledgerList = append(ledgerList, ledgerFrom(l))
	}
	return ledgerList, nil
}

func ledgerFrom(l Ledger) domain.Ledger {
	return domain.NewLedger(l.LedgerId, l.Amount, l.Title, l.Memo, l.Date, l.IsExcluded)
}
