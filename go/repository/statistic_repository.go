package repository

import (
	"context"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/domain"
	"github.com/yehey-1030/household-account-book/go/lib/transaction/gorm_tx"
	"github.com/yehey-1030/household-account-book/go/repository/database"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"gorm.io/gorm"
)

type StatisticRepository interface {
	Total(ctx context.Context, req domain.StatisticByTypeQuery) (int, error)
	TotalListByRootTag(ctx context.Context, req domain.TotalListByRootTagQuery) ([]domain.StatisticWithTag, error)
}

type statisticRepository struct {
	db *gorm.DB
}

func NewStatisticRepository(db *gorm.DB) StatisticRepository {
	return &statisticRepository{db: db}
}

type StatisticWithTag struct {
	TagId   int
	TagName string
	Total   int
}

func (s *statisticRepository) Total(ctx context.Context, req domain.StatisticByTypeQuery) (int, error) {
	db := gorm_tx.FromContextWithDefault(ctx, s.db)

	var total int
	result := db.Model(&database.Ledger{}).
		Select("COALESCE(sum(amount),0)").
		Where("ledger.date >= ? and ledger.date <= ?", req.StartDate, req.EndDate).
		Where("ledger.archivetype_id = ?", req.ArchiveTypeId).
		Find(&total)

	if result.Error != nil {
		return 0, fmt.Errorf("[%s] %w", ioutil.FuncName(), result.Error)
	}

	return total, nil
}

func (s *statisticRepository) TotalListByRootTag(ctx context.Context, req domain.TotalListByRootTagQuery) ([]domain.StatisticWithTag, error) {
	db := gorm_tx.FromContextWithDefault(ctx, s.db)

	var statistics []StatisticWithTag

	ledgerSubQuery := db.Model(&database.Ledger{}).Select("ledger_id,amount").Where("ledger.date >= ? and ledger.date <= ?", req.StartDate, req.EndDate)

	result := db.Model(&database.TagLedgerRelation{}).
		Table("`tag_ledger_relation` as tlr").
		Select("tag.tag_id, tag.tag_name, COALESCE(sum(l.amount),0) as total").
		Joins("left outer join (?) as l on tlr.ledger_id=l.ledger_id", ledgerSubQuery).
		Joins("left outer join tag on tlr.tag_id = tag.tag_id").
		Where("tag.archivetype_id = ? and tag.parent_id IS NULL ", req.ArchiveTypeId).
		Group("tlr.tag_id").
		Find(&statistics)

	if result.Error != nil {
		return nil, fmt.Errorf("[%s] %w", ioutil.FuncName(), result.Error)
	}

	var statisticDtoList []domain.StatisticWithTag

	for _, statistic := range statistics {
		statisticDtoList = append(statisticDtoList, domain.NewStatisticWithTag(statistic.Total, statistic.TagId, statistic.TagName))
	}
	return statisticDtoList, nil
}
