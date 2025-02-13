package beer

import (
	"context"

	"github.com/MikelSot/repository"

	"github.com/MikelSot/amaris-beer/insfrastructure/postgres/transaction"
	"github.com/MikelSot/amaris-beer/model"
)

type UseCase interface {
	Create(ctx context.Context, m *model.Beer) error
	Update(ctx context.Context, m model.Beer) error
	Delete(ctx context.Context, ID uint) error

	GetByID(ctx context.Context, ID uint) (model.Beer, error)
	GetAll(ctx context.Context) (model.Beers, error)
}

type Storage interface {
	Create(ctx context.Context, m *model.Beer) error
	Update(ctx context.Context, m model.Beer) error
	Delete(ctx context.Context, ID uint) error

	GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Beer, error)
	GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Beers, error)
}

type Tx interface {
	Begin(ctx context.Context) (transaction.Tx, error)
}
