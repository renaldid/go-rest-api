package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_rest_api/helper"
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
)

type OrderProductImpl struct {
}

func NewOrderProductRepository() OrdersProductRepository {
	return &OrderProductImpl{}
}

func (o OrderProductImpl) Save(ctx context.Context, tx *sql.Tx, ordersProduct domain.OrdersProduct) domain.OrdersProduct {
	SQL := "insert into orders_product (order_id, product_id, qty, amount, created_at, uploaded_at) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, ordersProduct.OrderId, ordersProduct.ProductId, ordersProduct.Qty, ordersProduct.Amount, ordersProduct.CreatedAt, ordersProduct.UploadedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	ordersProduct.Id = int(id)
	return ordersProduct
}

func (o OrderProductImpl) Update(ctx context.Context, tx *sql.Tx, ordersProduct domain.OrdersProduct) domain.OrdersProduct {
	SQL := "update orders_product set order_id = ?, product_id = ?, qty = ?, amount = ?, created_at = ?, uploaded_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, ordersProduct.OrderId, ordersProduct.ProductId, ordersProduct.Qty, ordersProduct.Amount, ordersProduct.CreatedAt, ordersProduct.UploadedAt, ordersProduct.Id)
	helper.PanicIfError(err)
	return ordersProduct
}

func (o OrderProductImpl) Delete(ctx context.Context, tx *sql.Tx, ordersProduct domain.OrdersProduct) {
	SQL := "delete from orders_product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, ordersProduct.Id)
	helper.PanicIfError(err)
}

func (o OrderProductImpl) FindById(ctx context.Context, tx *sql.Tx, ordersProductId int) (web.OrderProductResponse, error) {
	SQL := "select id, order_id, product_id, qty, amount, created_at, uploaded_at from orders_product where p.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, ordersProductId)
	helper.PanicIfError(err)
	defer rows.Close()

	orderProducts := web.OrderProductResponse{}
	if rows.Next() {
		err := rows.Scan(&orderProducts.Id, &orderProducts.OrderId, &orderProducts.ProductId, &orderProducts.Qty, &orderProducts.Amount, &orderProducts.CreatedAt, &orderProducts.UploadedAt)
		helper.PanicIfError(err)
		return orderProducts, nil
	} else {
		return orderProducts, errors.New("order product is not found")
	}
}

func (o OrderProductImpl) FindByAll(ctx context.Context, tx *sql.Tx) []web.OrderProductResponse {
	SQL := "select id, order_id, product_id, qty, amount, created_at, uploaded_at from orders_product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orderProduct []web.OrderProductResponse
	for rows.Next() {
		orderProducts := web.OrderProductResponse{}
		err := rows.Scan(&orderProducts.Id, &orderProducts.OrderId, &orderProducts.ProductId, &orderProducts.Qty, &orderProducts.Amount, &orderProducts.CreatedAt, &orderProducts.UploadedAt)
		helper.PanicIfError(err)
		orderProduct = append(orderProduct, orderProducts)
	}

	return orderProduct
}
