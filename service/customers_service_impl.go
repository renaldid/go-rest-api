package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"go_rest_api/exception"
	"go_rest_api/helper"
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
	"go_rest_api/repository"
)

type CustomersServiceImpl struct {
	CustomersRepository repository.CustomersRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewCustomersService(customersRepository repository.CustomersRepository, DB *sql.DB, Validate *validator.Validate) CustomersService {
	return &CustomersServiceImpl{
		CustomersRepository: customersRepository,
		DB:                  DB,
		Validate:            Validate,
	}
}

func (service *CustomersServiceImpl) Create(ctx context.Context, request web.CustomersCreateRequest) web.CustomersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers := domain.Customers{
		Name:        request.Name,
		Address:     request.Address,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}
	customers = service.CustomersRepository.Save(ctx, tx, customers)

	return helper.ToCustomersResponse(customers)
}

func (service *CustomersServiceImpl) Update(ctx context.Context, request web.CustomersUpdateRequest) web.CustomersResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers, err := service.CustomersRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	customers.Name = request.Name
	customers = service.CustomersRepository.Update(ctx, tx, customers)
	customers.Address = request.Address
	customers = service.CustomersRepository.Update(ctx, tx, customers)
	customers.Email = request.Email
	customers = service.CustomersRepository.Update(ctx, tx, customers)
	customers.PhoneNumber = request.PhoneNumber
	customers = service.CustomersRepository.Update(ctx, tx, customers)

	return helper.ToCustomersResponse(customers)
}

func (service *CustomersServiceImpl) Delete(ctx context.Context, customersId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers, err := service.CustomersRepository.FindById(ctx, tx, customersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.CustomersRepository.Delete(ctx, tx, customers)
}

func (service *CustomersServiceImpl) FindById(ctx context.Context, customersId int) web.CustomersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers, err := service.CustomersRepository.FindById(ctx, tx, customersId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCustomersResponse(customers)
}

func (service *CustomersServiceImpl) FindByAll(ctx context.Context) []web.CustomersResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := service.CustomersRepository.FindByAll(ctx, tx)

	return helper.ToCustomersResponses(customer)
}
