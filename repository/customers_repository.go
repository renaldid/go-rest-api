package repository

import (
	"context"
	"database/sql"
	"go_rest_api/model/domain"
)

type CustomersRepository interface {
	Save(ctx context.Context, tx *sql.Tx, customers domain.Customers) domain.Customers
	Update(ctx context.Context, tx *sql.Tx, customers domain.Customers) domain.Customers
	Delete(ctx context.Context, tx *sql.Tx, customers domain.Customers)
	FindById(ctx context.Context, tx *sql.Tx, customersId int) (domain.Customers, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Customers
}
