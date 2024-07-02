package database

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/lib/transaction/gorm_tx"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"gorm.io/gorm"
)

type Tag struct {
	TagId         int
	TagName       string
	ArchiveTypeId int `gorm:"column:archivetype_id"`
	ParentId      int
}

func (t Tag) TableName() string {
	return "tag"
}

type TagSearcher interface {
	ListByParent(ctx context.Context, parentId int) ([]domain.Tag, error)
}

type tagSearcher struct {
	db *gorm.DB
}

func NewTagSearcher(db *gorm.DB) TagSearcher {
	return &tagSearcher{db: db}
}

func (t *tagSearcher) ListByParent(ctx context.Context, parentId int) ([]domain.Tag, error) {
	db := gorm_tx.FromContextWithDefault(ctx, t.db)

	result := db.Model(&Tag{}).Select("tag.parent_id = ?", parentId)

	var tags []Tag
	result = result.Find(&tags)
	if result.Error != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), result.Error)
	}

	var tagList []domain.Tag
	for _, tag := range tags {
		tagList = append(tagList, tagFrom(tag))
	}
	return tagList, nil
}

func tagFrom(tag Tag) domain.Tag {
	return domain.NewTag(tag.TagId, tag.TagName, tag.ParentId, tag.ArchiveTypeId)
}
