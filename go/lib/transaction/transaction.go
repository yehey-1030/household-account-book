package transaction

import "context"

type Factory interface {
	Create(ctx context.Context) Transaction
}

type Transaction interface {
	Do(f func(ctx context.Context) error) error
}
