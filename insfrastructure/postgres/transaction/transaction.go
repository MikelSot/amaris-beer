package transaction

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Transaction struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Transaction {
	return Transaction{db}
}

func (t Transaction) Begin(ctx context.Context) (Tx, error) {
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return Tx{}, err
	}

	return new(tx), nil
}
