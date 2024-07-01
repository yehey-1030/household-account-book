package domain

import "time"

type Ledger interface {
	Id() int
	Title() string
	Memo() string
	Date() *time.Time
	Amount() int
	IsExcluded() bool
	ArchiveTypeId() int
}

type ledger struct {
	ledgerId      int
	amount        int
	date          *time.Time
	title         string
	memo          string
	isExcluded    bool
	archiveTypeId int
}

func (l *ledger) Id() int {
	return l.ledgerId
}

func (l *ledger) Title() string {
	return l.title
}

func (l *ledger) Memo() string {
	return l.memo
}

func (l *ledger) Date() *time.Time {
	return l.date
}

func (l *ledger) Amount() int {
	return l.amount
}

func (l *ledger) IsExcluded() bool {
	return l.isExcluded
}

func (l *ledger) ArchiveTypeId() int {
	return l.archiveTypeId
}

func NewLedger(id, amount int, title, memo string, date *time.Time, isExcluded bool, archiveTypeId int) Ledger {
	return &ledger{ledgerId: id, amount: amount, title: title, memo: memo, date: date, isExcluded: isExcluded, archiveTypeId: archiveTypeId}
}

type LedgerPagingQuery struct {
	StartDate     time.Time
	EndDate       time.Time
	ArchiveTypeId int
	TagId         int
}
