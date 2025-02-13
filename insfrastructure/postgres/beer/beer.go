package beer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/MikelSot/amaris-beer/model"
	"github.com/MikelSot/repository"
)

var _fieldInserts = []string{
	"name",
	"price",
	"description",
}

var _fieldsSelect = []string{
	"id",
	"created_at",
	"updated_at",
}

const _table = "beers"

var (
	_psqlInsert = repository.BuildSQLInsertNoID(_table, _fieldInserts)
	_psqlUpdate = repository.BuildSQLUpdateByID(_table, _fieldInserts)
	_psqlDelete = "DELETE FROM " + _table + " WHERE id = $1"

	_psqlGetAll = repository.BuildSQLSelectFields(_table, append(_fieldInserts, _fieldsSelect...))
)

type Beer struct {
	db model.PgxPool
}

func New(db model.PgxPool) Beer {
	return Beer{db}
}

func (b Beer) Create(ctx context.Context, m *model.Beer) error {
	err := b.db.QueryRow(
		ctx,
		_psqlInsert,
		m.Name,
		m.Price,
		repository.StringToNull(m.Description),
	).Scan(&m.ID, &m.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (b Beer) Update(ctx context.Context, m model.Beer) error {
	_, err := b.db.Exec(
		ctx,
		_psqlUpdate,
		m.Name,
		m.Price,
		repository.StringToNull(m.Description),
		m.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (b Beer) Delete(ctx context.Context, ID uint) error {
	_, err := b.db.Exec(
		ctx,
		_psqlDelete,
		ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (b Beer) GetWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Beer, error) {
	query, args := repository.BuildQueryAndArgs(_psqlGetAll, specification)

	m, err := b.scanRow(b.db.QueryRow(ctx, query, args...))
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Beer{}, nil
	}
	if err != nil {
		return model.Beer{}, err
	}

	return m, nil
}

func (b Beer) GetAllWhere(ctx context.Context, specification repository.FieldsSpecification) (model.Beers, error) {
	query, args := repository.BuildQueryArgsAndPagination(_psqlGetAll, specification)

	rows, err := b.db.Query(ctx, query, args...)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	ms := model.Beers{}
	for rows.Next() {
		m, err := b.scanRow(rows)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func (b Beer) scanRow(row pgx.Row) (model.Beer, error) {
	m := model.Beer{}

	descriptionNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := row.Scan(
		&m.Name,
		&m.Price,
		&descriptionNull,
		&m.ID,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return model.Beer{}, err
	}

	m.Description = descriptionNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
