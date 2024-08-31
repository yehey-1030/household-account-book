package database

import (
	"context"
	"database/sql"
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
	ParentId      sql.NullInt64
}

func (t Tag) TableName() string {
	return "tag"
}

type TagLedgerRelation struct {
	RelationId int
	TagId      int
	LedgerId   int
}

func (r TagLedgerRelation) TableName() string {
	return "tag_ledger_relation"
}

type TagSearcher interface {
	ListByParent(ctx context.Context, parentId int) ([]domain.Tag, error)
	ListRootTag(ctx context.Context, archiveTypeId int) ([]domain.Tag, error)
	Create(ctx context.Context, tag domain.Tag) (domain.Tag, error)
}

type tagSearcher struct {
	db *gorm.DB
}

func NewTagSearcher(db *gorm.DB) TagSearcher {
	return &tagSearcher{db: db}
}

func (t *tagSearcher) ListByParent(ctx context.Context, parentId int) ([]domain.Tag, error) {
	db := gorm_tx.FromContextWithDefault(ctx, t.db)

	result := db.Model(&Tag{}).
		Where("tag.parent_id = ?", parentId)

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

func (t *tagSearcher) ListRootTag(ctx context.Context, archiveTypeId int) ([]domain.Tag, error) {
	db := gorm_tx.FromContextWithDefault(ctx, t.db)

	result := db.Model(&Tag{}).
		Where("tag.archivetype_id = ?", archiveTypeId).
		Where("tag.parent_id IS NULL")

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

func (t *tagSearcher) Create(ctx context.Context, tag domain.Tag) (domain.Tag, error) {
	db := gorm_tx.FromContextWithDefault(ctx, t.db)

	tagDto := TagDtoFrom(tag)

	result := db.Create(&tagDto)
	if result.Error != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), result.Error)
	}

	created := tagFrom(tagDto)
	return created, nil
}

func tagFrom(tag Tag) domain.Tag {
	return domain.NewTag(tag.TagId, tag.TagName, ioutil.NullIntToInt(tag.ParentId), tag.ArchiveTypeId)
}

func TagDtoFrom(tag domain.Tag) Tag {
	return Tag{
		TagId:         0,
		TagName:       tag.Name(),
		ArchiveTypeId: tag.ArchiveTypeId(),
		ParentId:      ioutil.IntToNullInt(tag.ParentId()),
	}
}
