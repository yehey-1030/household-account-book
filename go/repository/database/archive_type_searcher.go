package database

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/lib/transaction/gorm_tx"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"gorm.io/gorm"
)

type ArchiveType struct {
	ArchiveTypeId int `gorm:"column:archivetype_id"`
	TypeName      string
}

type ArchiveTypeSearcher interface {
	List(ctx context.Context) ([]domain.ArchiveType, error)
	Get(ctx context.Context, id int) (domain.ArchiveType, error)
}

type archiveTypeSearcher struct {
	db *gorm.DB
}

func NewArchiveTypeSearcher(db *gorm.DB) ArchiveTypeSearcher {
	return &archiveTypeSearcher{db: db}
}

func (a *archiveTypeSearcher) List(ctx context.Context) ([]domain.ArchiveType, error) {
	db := gorm_tx.FromContextWithDefault(ctx, a.db)

	var archiveTypes []ArchiveType
	result := db.Model(&ArchiveType{}).Find(archiveTypes)
	if result.Error != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), result.Error)
	}

	var archiveTypeList []domain.ArchiveType
	for _, a := range archiveTypes {
		archiveTypeList = append(archiveTypeList, archiveTypeFrom(a))
	}

	return archiveTypeList, nil
}

func (a *archiveTypeSearcher) Get(ctx context.Context, id int) (domain.ArchiveType, error) {
	db := gorm_tx.FromContextWithDefault(ctx, a.db)

	var archiveType ArchiveType
	result := db.Model(&ArchiveType{}).Where("archivetype_id = ?", id).Find(&archiveType)
	if result.Error != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), result.Error)
	}

	return archiveTypeFrom(archiveType), nil
}

func archiveTypeFrom(a ArchiveType) domain.ArchiveType {
	return domain.NewArchiveType(a.ArchiveTypeId, a.TypeName)
}
