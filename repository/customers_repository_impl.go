package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_rest_api/helper"
	"go_rest_api/model/domain"
)

type CustomersRepositoryImpl struct {
}

func NewCustomersRepository() CustomersRepository {
	return &CustomersRepositoryImpl{}
}

func (c CustomersRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, customers domain.Customers) domain.Customers {
	SQL := "insert into customers_model(name, address, email, phoneNumber ) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, customers.Name, customers.Address, customers.Email, customers.PhoneNumber)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	customers.Id = int(id)
	return customers
}

func (c CustomersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customers domain.Customers) domain.Customers {
	//UPDATE `customers_model` SET `address` = 'jalan belum ada', `email` = 'renaldi@gmail.com', `phoneNumber` = '081233244335' WHERE `customers_model`.`id` = 1;
	SQL := "update customers_model set name = ?, address = ?, email = ?, phoneNumber = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customers.Name, customers.Address, customers.Email, customers.PhoneNumber, customers.Id)
	helper.PanicIfError(err)

	return customers
}

func (c CustomersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customers domain.Customers) {
	SQL := "delete from customers_model where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customers.Id)
	helper.PanicIfError(err)
}

func (c CustomersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customersId int) (domain.Customers, error) {
	SQL := "select id, name, address, email, phoneNumber from customers_model where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, customersId)
	helper.PanicIfError(err)
	defer rows.Close()

	customers := domain.Customers{}
	if rows.Next() {
		err := rows.Scan(&customers.Id, &customers.Name, &customers.Address, &customers.Email, &customers.PhoneNumber)
		helper.PanicIfError(err)
		return customers, nil
	} else {
		return customers, errors.New("customer is not found")
	}
}

func (c CustomersRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Customers {
	SQL := "select id, name, address, email, phoneNumber from customers_model"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var customer []domain.Customers
	for rows.Next() {
		customers := domain.Customers{}
		err := rows.Scan(&customers.Id, &customers.Name, &customers.Address, &customers.Email, &customers.PhoneNumber)
		helper.PanicIfError(err)
		customer = append(customer, customers)
	}
	return customer
}
