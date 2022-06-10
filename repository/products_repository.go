package repository

import (
	"context"
	"database/sql"
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
)

type ProductsRepository interface {
	Save(ctx context.Context, tx *sql.Tx, products domain.Products) domain.Products
	Update(ctx context.Context, tx *sql.Tx, products domain.Products) domain.Products
	Delete(ctx context.Context, tx *sql.Tx, products domain.Products)
	FindById(ctx context.Context, tx *sql.Tx, productsId int) (web.ProductsResponse, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []web.ProductsResponse
}
