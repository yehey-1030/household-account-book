package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

type LedgerSearcherSuite struct {
	suite.Suite
	ledgerSearcher database.LedgerSearcher
	db             *gorm.DB
	date           time.Time
}

func TestLedgerSearcherSuite(t *testing.T) {
	suite.Run(t, new(LedgerSearcherSuite))
}

func (s *LedgerSearcherSuite) SetupTest() {
	logConfig := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      logger.Info,
		Colorful:      true,
	})

	var err error
	s.db, err = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{Logger: logConfig})
	if err != nil {
		panic("failed to connect database")
	}

	_ = s.db.AutoMigrate(&database.Ledger{})
	s.date = time.Now().UTC()

	ledger1 := database.Ledger{
		LedgerId: 1,
		Title:    "l-1",
		Memo:     "l-1",
		Amount:   10,
		Date:     &s.date,
	}
	ledger2 := database.Ledger{
		LedgerId: 2,
		Title:    "l-2",
		Memo:     "l-2",
		Amount:   20,
		Date:     &s.date,
	}
	s.db.Create(ledger1)
	s.db.Create(ledger2)

	s.ledgerSearcher = database.NewLedgerSearcher(s.db)
}
func (s *LedgerSearcherSuite) TearDownSuite() {
	sqlDB, _ := s.db.DB()
	_ = sqlDB.Close()
}

func (s *LedgerSearcherSuite) TestList() {
	ctx := context.Background()

	query := domain.LedgerPagingQuery{
		StartDate:     s.date.AddDate(-1, 0, 0),
		EndDate:       s.date.AddDate(0, 1, 0),
		TagId:         1,
		ArchiveTypeId: 1,
	}

	list, err := s.ledgerSearcher.List(ctx, query)
	s.Nil(err)
	s.Len(list, 2)
}
