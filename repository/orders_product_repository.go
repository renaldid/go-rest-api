package repository

import (
	"context"
	"database/sql"
	"go_rest_api/model/domain"
)

type OrdersProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, ordersProduct domain.OrdersProduct) domain.OrdersProduct
	Update(ctx context.Context, tx *sql.Tx, ordersProduct domain.OrdersProduct) domain.OrdersProduct
	Delete(ctx context.Context, tx *sql.Tx, ordersProduct domain.OrdersProduct)
	FindById(ctx context.Context, tx *sql.Tx, ordersProductId int) (domain.OrdersProduct, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.OrdersProduct
}
