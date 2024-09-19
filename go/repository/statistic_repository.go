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
}

type statisticRepository struct {
	db *gorm.DB
}

func NewStatisticRepository(db *gorm.DB) StatisticRepository {
	return &statisticRepository{db: db}
}

func (s *statisticRepository) Total(ctx context.Context, req domain.StatisticByTypeQuery) (int, error) {
	db := gorm_tx.FromContextWithDefault(ctx, s.db)

	var total int
	result := db.Model(&database.Ledger{}).
		Select("sum(ledger.amount)").
		Where("ledger.date >= ? and ledger.date <= ?", req.StartDate, req.EndDate).
		Where("ledger.archivetype_id = ?", req.ArchiveTypeId).
		Find(&total)

	if result.Error != nil {
		return 0, fmt.Errorf("[%s] %w", ioutil.FuncName(), result.Error)
	}

	return total, nil
}
