package transaction

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Tx struct {
	tx pgx.Tx
}

func new(tx pgx.Tx) Tx {
	return Tx{tx}
}

func (t Tx) Commit() error {
	return t.tx.Commit(context.Background())

}

func (t Tx) Rollback() error {
	return t.tx.Rollback(context.Background())
}
