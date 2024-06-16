package gorm_tx

import (
	"context"
	"github.com/yehey-1030/household-account-book/go/lib/transaction"
	"gorm.io/gorm"
)

//https://gorm.io/docs/transactions.html

type factory struct {
	db *gorm.DB
}

func NewFactory(db *gorm.DB) *factory {
	return &factory{db: db}
}

func (f *factory) Create(ctx context.Context) transaction.Transaction {
	tx := f.db.Begin()
	newCtx := context.WithValue(ctx, dbSessionKey{}, tx)
	return &GormTransaction{ctx: newCtx, tx: tx}
}

type GormTransaction struct {
	ctx context.Context
	tx  *gorm.DB
}

type dbSessionKey struct{}

func (t *GormTransaction) Do(f func(ctx context.Context) error) error {
	err := f(t.ctx)
	if err != nil {
		t.tx.Rollback()
		return err
	}
	t.tx.Commit()
	return nil
}

func FromContext(ctx context.Context) (*gorm.DB, bool) {
	session, ok := ctx.Value(dbSessionKey{}).(*gorm.DB)
	return session, ok
}

func FromContextWithDefault(ctx context.Context, db *gorm.DB) *gorm.DB {
	session, ok := FromContext(ctx)
	if !ok {
		return db
	}
	return session
}
