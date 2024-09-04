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
	tagSearcher    database.TagSearcher
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

	_ = s.db.AutoMigrate(&database.Tag{}, &database.Ledger{}, &database.TagLedgerRelation{}, &database.ArchiveType{})
	s.date = time.Now().UTC()

	archiveType := database.ArchiveType{
		ArchiveTypeId: 0,
		TypeName:      "at-1",
	}
	s.db.Create(&archiveType)
	ledger1 := database.Ledger{
		LedgerId:      1,
		Title:         "l-1",
		Memo:          "l-1",
		Amount:        10,
		Date:          &s.date,
		ArchiveTypeId: 1,
	}
	ledger2 := database.Ledger{
		LedgerId:      2,
		Title:         "l-2",
		Memo:          "l-2",
		Amount:        20,
		Date:          &s.date,
		ArchiveTypeId: 1,
	}
	s.db.Omit("Tags").Omit("ArchiveType").Create(&ledger1)
	s.db.Omit("Tags").Omit("ArchiveType").Create(&ledger2)

	tag1 := database.Tag{
		TagId:         1,
		TagName:       "tag1",
		ArchiveTypeId: 3,
	}
	tag2 := database.Tag{
		TagId:         2,
		TagName:       "tag2",
		ArchiveTypeId: 3,
	}
	s.db.Create(&tag1)
	s.db.Create(&tag2)

	tagLedger1 := database.TagLedgerRelation{
		TagId:    1,
		LedgerId: 1,
	}
	tagLedger2 := database.TagLedgerRelation{
		TagId:    2,
		LedgerId: 1,
	}
	s.db.Create(&tagLedger1)
	s.db.Create(&tagLedger2)

	s.ledgerSearcher = database.NewLedgerSearcher(s.db)
	s.tagSearcher = database.NewTagSearcher(s.db)
}
func (s *LedgerSearcherSuite) TearDownSuite() {
	sqlDB, _ := s.db.DB()
	_ = sqlDB.Close()
}

func (s *LedgerSearcherSuite) TestList() {
	ctx := context.Background()

	query := domain.LedgerPagingQuery{
		StartDate:     s.date.AddDate(-1, 0, 0),
		EndDate:       s.date.AddDate(0, 0, 1),
		TagId:         1,
		ArchiveTypeId: 1,
	}

	list, err := s.ledgerSearcher.List(ctx, query)
	s.Nil(err)
	s.Len(list, 2)
}

func (s *LedgerSearcherSuite) TestCreate() {
	ctx := context.Background()

	today := s.date.AddDate(0, 1, 0)
	archiveType := domain.NewArchiveType(3, "at-3")
	tags := []domain.Tag{
		domain.NewTag(3, "new_tag", 0, 3),
		domain.NewTag(1, "tag1", 0, 3),
	}
	toCreate := domain.NewLedger(0, 10000, "test", "test", &today, false, archiveType, tags)

	created, err := s.ledgerSearcher.Create(ctx, toCreate)
	s.Nil(err)
	s.Equal(created.Amount(), 10000)

	t, err := s.tagSearcher.ListRootTag(ctx, 3)
	s.Nil(err)
	s.Len(t, 3)
}
