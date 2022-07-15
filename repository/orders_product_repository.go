package repository

import (
	"context"
	"database/sql"
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
)

type OrdersProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, ordersProduct domain.OrdersProduct) domain.OrdersProduct
	Update(ctx context.Context, tx *sql.Tx, ordersProduct domain.OrdersProduct) domain.OrdersProduct
	Delete(ctx context.Context, tx *sql.Tx, ordersProduct domain.OrdersProduct)
	FindById(ctx context.Context, tx *sql.Tx, ordersProductId int) (web.OrderProductResponse, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []web.OrderProductResponse
}
