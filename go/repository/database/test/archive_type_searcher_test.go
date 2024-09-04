package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/yehey-1030/household-account-book/go/repository/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

type ArchiveTypeSearcherSuite struct {
	suite.Suite
	archiveTypeSearcher database.ArchiveTypeSearcher
	db                  *gorm.DB
}

func TestArchiveTypeSearcherSuite(t *testing.T) {
	suite.Run(t, new(ArchiveTypeSearcherSuite))
}

func (s *ArchiveTypeSearcherSuite) SetupTest() {
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

	_ = s.db.AutoMigrate(&database.ArchiveType{})

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

	s.archiveTypeSearcher = database.NewArchiveTypeSearcher(s.db)
}

func (s *ArchiveTypeSearcherSuite) TestList() {
	ctx := context.Background()

	list, err := s.archiveTypeSearcher.List(ctx)
	s.Nil(err)
	s.Len(list, 2)
}

func (s *ArchiveTypeSearcherSuite) TestGetById() {
	ctx := context.Background()

	archiveType, err := s.archiveTypeSearcher.GetById(ctx, 1)
	s.Nil(err)
	s.Equal("income", archiveType.Name())
}
