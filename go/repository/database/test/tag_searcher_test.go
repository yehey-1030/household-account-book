package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/repository/database"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

type TagSearcherSuite struct {
	suite.Suite
	tagSearcher database.TagSearcher
	db          *gorm.DB
}

func TestTagSearcherSuite(t *testing.T) {
	suite.Run(t, new(TagSearcherSuite))
}

func (s *TagSearcherSuite) SetupTest() {
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

	_ = s.db.AutoMigrate(&database.Tag{})

	tag1 := database.Tag{
		TagId:         1,
		TagName:       "parent",
		ParentId:      ioutil.IntToNullInt(0),
		ArchiveTypeId: 1,
	}
	tag2 := database.Tag{
		TagId:         2,
		TagName:       "child1",
		ParentId:      ioutil.IntToNullInt(1),
		ArchiveTypeId: 1,
	}
	tag3 := database.Tag{
		TagId:         3,
		TagName:       "child2",
		ParentId:      ioutil.IntToNullInt(1),
		ArchiveTypeId: 1,
	}
	tag4 := database.Tag{
		TagId:         4,
		TagName:       "parent2",
		ParentId:      ioutil.IntToNullInt(0),
		ArchiveTypeId: 1,
	}

	s.db.Create(tag1)
	s.db.Create(tag2)
	s.db.Create(tag3)
	s.db.Create(tag4)

	s.tagSearcher = database.NewTagSearcher(s.db)
}
func (s *TagSearcherSuite) TearDownSuite() {
	sqlDB, _ := s.db.DB()
	_ = sqlDB.Close()
}

func (s *TagSearcherSuite) TestListByParent() {
	ctx := context.Background()

	list, err := s.tagSearcher.ListByParent(ctx, 1)
	s.Nil(err)
	s.Len(list, 2)
	s.Equal(list[0].ParentId(), 1)
}

func (s *TagSearcherSuite) TestListRootTag() {
	ctx := context.Background()

	list, err := s.tagSearcher.ListRootTag(ctx, 1)
	s.Nil(err)
	s.Len(list, 2)
}

func (s *TagSearcherSuite) TestCreate() {
	ctx := context.Background()

	toCreate := domain.NewTag(0, "newTag", 3, 1)

	created, err := s.tagSearcher.Create(ctx, toCreate)
	s.Nil(err)
	s.Equal(created.ParentId(), 3)
}
