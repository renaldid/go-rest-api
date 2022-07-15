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

type OrderProductsServiceImpl struct {
	OrderProductsRepository repository.OrdersProductRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewOrderProductsService(OrderProductsRepository repository.OrdersProductRepository, DB *sql.DB, Validate *validator.Validate) OrderProductsService {
	return &OrderProductsServiceImpl{
		OrderProductsRepository: OrderProductsRepository,
		DB:                      DB,
		Validate:                Validate,
	}
}

func (service *OrderProductsServiceImpl) Create(ctx context.Context, request web.OrderProductCreateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderProduct := domain.OrdersProduct{
		OrderId:    request.OrderId,
		ProductId:  request.ProductId,
		Qty:        request.Qty,
		Amount:     request.Amount,
		CreatedAt:  request.CreatedAt,
		UploadedAt: request.UploadedAt,
	}

	orderProduct = service.OrderProductsRepository.Save(ctx, tx, orderProduct)
	return helper.ToOrderProductResponse(orderProduct)
}

func (service *OrderProductsServiceImpl) Update(ctx context.Context, request web.OrderProductUpdateRequest) web.OrderProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productsResponse, err := service.OrderProductsRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	product := helper.ToOrderProduct(productsResponse)
	product.OrderId = request.OrderId
	product.ProductId = request.ProductId
	product.Qty = request.Qty
	product.Amount = request.Amount
	product.CreatedAt = request.CreatedAt
	product.UploadedAt = request.UploadedAt
	product = service.OrderProductsRepository.Update(ctx, tx, product)

	return helper.ToOrderProductResponse(product)
}

func (service *OrderProductsServiceImpl) Delete(ctx context.Context, orderProductId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productResponse, err := service.OrderProductsRepository.FindById(ctx, tx, orderProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	product := helper.ToOrderProduct(productResponse)
	service.OrderProductsRepository.Delete(ctx, tx, product)
}

func (service *OrderProductsServiceImpl) FindById(ctx context.Context, orderProductId int) web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products, err := service.OrderProductsRepository.FindById(ctx, tx, orderProductId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return products
}

func (service *OrderProductsServiceImpl) FindByAll(ctx context.Context) []web.OrderProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.OrderProductsRepository.FindByAll(ctx, tx)

	return products
}
