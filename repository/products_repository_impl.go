package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go_rest_api/helper"
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
)

type ProductsRepositoryImpl struct {
}

func NewProductsRepository() ProductsRepository {
	return &ProductsRepositoryImpl{}
}

func (p ProductsRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, products domain.Products) domain.Products {
	//ini membuat logger pembuka
	logrus.Info("product repository start")

	SQL := "insert into products (name, price, category_id) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, products.Name, products.Price, products.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	products.Id = int(id)
	//ini membuat logger penutup
	logrus.Info("product repository ended")

	return products
}

func (p ProductsRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, products domain.Products) domain.Products {
	//ini membuat logger pembuka
	logrus.Info("product repository start")

	SQL := "update products set name = ?, price = ?, category_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, products.Name, products.Price, products.CategoryId, products.Id)
	helper.PanicIfError(err)
	//ini membuat logger penutup
	logrus.Info("product repository ended")

	return products
}

func (p ProductsRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, products domain.Products) {
	//ini membuat logger pembuka
	logrus.Info("product repository start")

	SQL := "delete from products where id = ?"
	_, err := tx.ExecContext(ctx, SQL, products.Id)
	helper.PanicIfError(err)
	//ini membuat logger penutup
	logrus.Info("product repository ended")
}

func (p ProductsRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productsId int) (web.ProductsResponse, error) {
	//ini membuat logger pembuka
	logrus.Info("product repository start")

	SQL := "select p.id, p.name, p.price, p.category_id, c.name from products p inner join category c on p.category_id=c.id where p.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productsId)
	helper.PanicIfError(err)
	defer rows.Close()

	products := web.ProductsResponse{}
	if rows.Next() {
		err := rows.Scan(&products.Id, &products.Name, &products.Price, &products.CategoryId, &products.CategoryName)
		helper.PanicIfError(err)
		//ini membuat logger penutup
		logrus.Info("product repository ended")

		return products, nil
	} else {
		//ini membuat logger penutup
		logrus.Info("product repository ended")

		return products, errors.New("product is not found")
	}
}

func (p ProductsRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []web.ProductsResponse {
	//ini membuat logger pembuka
	logrus.Info("product repository start")

	SQL := "select p.id, p.name, p.price, p.category_id, c.name from products p inner join category c on p.category_id=c.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var product []web.ProductsResponse
	for rows.Next() {
		products := web.ProductsResponse{}
		err := rows.Scan(&products.Id, &products.Name, &products.Price, &products.CategoryId, &products.CategoryName)
		helper.PanicIfError(err)
		product = append(product, products)
	}
	//ini membuat logger penutup
	logrus.Info("product repository ended")

	return product
}
