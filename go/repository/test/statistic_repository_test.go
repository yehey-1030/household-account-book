package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository"
	"github.com/yehey-1030/household-account-book/go/repository/database"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"github.com/yehey-1030/household-account-book/go/util/timeutil"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

type StatisticRepositorySuite struct {
	suite.Suite
	statisticRepository repository.StatisticRepository
	db                  *gorm.DB
	date                time.Time
}

func TestStatisticRepositorySuite(t *testing.T) {
	suite.Run(t, new(StatisticRepositorySuite))
}

func (s *StatisticRepositorySuite) SetupTest() {
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

	s.date = time.Now().UTC()

	_ = s.db.AutoMigrate(&database.ArchiveType{}, &database.Ledger{}, &database.Tag{})

	type1 := database.ArchiveType{
		ArchiveTypeId: 1,
		TypeName:      "income",
	}
	type2 := database.ArchiveType{
		ArchiveTypeId: 2,
		TypeName:      "fund",
	}
	s.db.Create(&type1)
	s.db.Create(&type2)

	ledger1 := database.Ledger{
		LedgerId:      0,
		Title:         "l-1",
		Memo:          "l-1",
		Amount:        10,
		Date:          &s.date,
		ArchiveTypeId: 1,
		Tags: []database.Tag{
			database.Tag{
				TagId:         1,
				TagName:       "tag1",
				ParentId:      ioutil.IntToNullInt(0),
				ArchiveTypeId: 1,
			},
			database.Tag{
				TagId:         2,
				TagName:       "tag2",
				ParentId:      ioutil.IntToNullInt(0),
				ArchiveTypeId: 1,
			},
		},
	}
	ledger2 := database.Ledger{
		LedgerId:      0,
		Title:         "l-2",
		Memo:          "l-2",
		Amount:        20,
		Date:          timeutil.StringToDate(s.date.Format(time.DateOnly)),
		ArchiveTypeId: 1,
		Tags: []database.Tag{
			database.Tag{
				TagId:         1,
				TagName:       "tag1",
				ParentId:      ioutil.IntToNullInt(0),
				ArchiveTypeId: 1,
			},
		},
	}
	ledger3 := database.Ledger{
		LedgerId:      0,
		Title:         "l-3",
		Memo:          "l-3",
		Amount:        40,
		Date:          timeutil.StringToDate(s.date.Format(time.DateOnly)),
		ArchiveTypeId: 2,
		Tags: []database.Tag{
			database.Tag{
				TagId:         3,
				TagName:       "tag3",
				ParentId:      ioutil.IntToNullInt(0),
				ArchiveTypeId: 1,
			},
		},
	}
	s.db.Create(&ledger1)
	s.db.Create(&ledger2)
	s.db.Create(&ledger3)

	s.statisticRepository = repository.NewStatisticRepository(s.db)
}

func (s *StatisticRepositorySuite) TearDownSuite() {
	sqlDB, _ := s.db.DB()
	_ = sqlDB.Close()
}

func (s *StatisticRepositorySuite) TestTotal() {
	ctx := context.Background()

	req := domain.StatisticByTypeQuery{
		ArchiveTypeId: 1,
		StartDate:     s.date.AddDate(-1, 0, 0).Format(time.DateOnly),
		EndDate:       s.date.AddDate(0, 0, 1).Format(time.DateOnly),
	}

	total, err := s.statisticRepository.Total(ctx, req)
	s.Nil(err)
	s.Equal(total, 30)
}

func (s *StatisticRepositorySuite) TestTotalListByRootTag() {
	ctx := context.Background()

	req := domain.TotalListByRootTagQuery{
		ArchiveTypeId: 1,
		StartDate:     s.date.AddDate(0, 0, -1).Format(time.DateOnly),
		EndDate:       s.date.AddDate(0, 0, 1).Format(time.DateOnly),
	}

	statistics, err := s.statisticRepository.TotalListByRootTag(ctx, req)
	s.Nil(err)
	s.Len(statistics, 3)
}
