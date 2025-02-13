package beer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/MikelSot/repository"

	"github.com/MikelSot/amaris-beer/model"
)

type Beer struct {
	storage Storage
	tx      Tx
}

func New(s Storage, tx Tx) Beer {
	return Beer{s, tx}
}

func (b Beer) Create(ctx context.Context, m *model.Beer) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("beer: %w", err)
	}

	if err := b.validateUniqueName(ctx, *m); err != nil {
		return err
	}

	tx, err := b.tx.Begin(ctx)
	if err != nil {
		return fmt.Errorf("beer.tx.Begin(): %w", err)
	}

	if err := b.storage.Create(ctx, m); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("beer: %s, %w", rollbackErr, err)
		}

		return b.errorConstraint(fmt.Errorf("beer.storage.Create(): %w", err))
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("beer: could not commit transaction: %w", err)
	}

	return nil
}

func (b Beer) Update(ctx context.Context, m model.Beer) error {
	if err := model.ValidateStructNil(m); err != nil {
		return fmt.Errorf("beer: %w", err)
	}

	if !m.HasID() {
		return model.ErrInvalidID
	}

	tx, err := b.tx.Begin(ctx)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	if err := b.storage.Update(ctx, m); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("update failed: %w, rollback failed: %v", err, rollbackErr)
		}
		return b.errorConstraint(err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

func (b Beer) Delete(ctx context.Context, ID uint) error {
	err := b.storage.Delete(ctx, ID)
	if err != nil {
		return b.errorConstraint(err)
	}

	return nil
}

func (b Beer) GetByID(ctx context.Context, ID uint) (model.Beer, error) {
	if ID == 0 {
		return model.Beer{}, model.ErrInvalidID
	}

	beer, err := b.storage.GetWhere(ctx, repository.FieldsSpecification{Filters: repository.Fields{{Name: "id", Value: ID}}})
	if err != nil {
		return model.Beer{}, fmt.Errorf("beer: %w", err)
	}

	return beer, nil
}

func (b Beer) GetAll(ctx context.Context) (model.Beers, error) {
	ms, err := b.storage.GetAllWhere(ctx, repository.FieldsSpecification{})
	if err != nil {
		return nil, fmt.Errorf("beer.storage.GetAllWhere(): %w", err)
	}

	return ms, nil
}

func (b Beer) validateUniqueName(ctx context.Context, m model.Beer) error {
	customErr := model.NewError()

	if m.Name == "" {
		customErr.SetStatusHTTP(http.StatusBadRequest)
		customErr.Fields.Add(model.ErrorDetail{
			Field:       "name",
			Issue:       model.IssueViolatedValidation,
			Description: "name is required",
		})
		customErr.SetAPIMessage("¡Upps! El campo name es requerido.")

		return customErr
	}

	beer, err := b.getByName(ctx, m.Name)
	if err != nil {
		return err
	}

	if beer.HasID() {
		customErr.SetStatusHTTP(http.StatusConflict)
		customErr.Fields.Add(model.ErrorDetail{
			Field:       "name",
			Issue:       model.IssueViolatedValidation,
			Description: "name already exists",
		})
		customErr.SetAPIMessage("¡Upps! El campo name ya existe.")

		return customErr
	}

	return nil
}

func (b Beer) getByName(ctx context.Context, name string) (model.Beer, error) {
	beer, err := b.storage.GetWhere(ctx, repository.FieldsSpecification{Filters: repository.Fields{{Name: "name", Value: name}}})
	if err != nil {
		return model.Beer{}, fmt.Errorf("beer: %w", err)
	}

	return beer, nil
}

func (b Beer) errorConstraint(err error) error {
	e := model.NewError()
	e.SetError(err)

	switch err {
	default:
		return err
	}
}
