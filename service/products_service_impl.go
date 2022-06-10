package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"go_rest_api/exception"
	"go_rest_api/helper"
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
	"go_rest_api/repository"
)

type ProductsServiceImpl struct {
	ProductsRepository repository.ProductsRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewProductsService(productsRepository repository.ProductsRepository, DB *sql.DB, Validate *validator.Validate) ProductsService {
	return &ProductsServiceImpl{
		ProductsRepository: productsRepository,
		DB:                 DB,
		Validate:           Validate,
	}
}

func (service *ProductsServiceImpl) Create(ctx context.Context, request web.ProductsCreateRequest) web.ProductsResponse {
	//ini membuat logger pembuka
	logrus.Info("product service start")

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := domain.Products{
		Name:       request.Name,
		Price:      request.Price,
		CategoryId: request.CategoryId,
	}
	products = service.ProductsRepository.Save(ctx, tx, products)
	//ini membuat logger penutup
	logrus.Info("product service ended")

	return helper.ToProductsResponse(products)

}

func (service *ProductsServiceImpl) Update(ctx context.Context, request web.ProductsUpdateRequest) web.ProductsResponse {
	//ini membuat logger pembuka
	logrus.Info("product service start")

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productsResponse, err := service.ProductsRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	product := helper.ToProducts(productsResponse)
	product.Name = request.Name
	product.Price = request.Price
	product.CategoryId = request.CategoryId
	product = service.ProductsRepository.Update(ctx, tx, product)

	//ini membuat logger penutup
	logrus.Info("product service ended")

	return helper.ToProductsResponse(product)
}

func (service *ProductsServiceImpl) Delete(ctx context.Context, productsId int) {

	//ini membuat logger pembuka
	logrus.Info("product service start")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productResponse, err := service.ProductsRepository.FindById(ctx, tx, productsId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//ini membuat logger penutup
	logrus.Info("product service ended")

	product := helper.ToProducts(productResponse)
	service.ProductsRepository.Delete(ctx, tx, product)
}

func (service *ProductsServiceImpl) FindById(ctx context.Context, productsId int) web.ProductsResponse {
	//ini membuat logger pembuka
	logrus.Info("product service start")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products, err := service.ProductsRepository.FindById(ctx, tx, productsId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//ini membuat logger penutup
	logrus.Info("product service ended")

	return products
}

func (service *ProductsServiceImpl) FindByAll(ctx context.Context) []web.ProductsResponse {
	//ini membuat logger pembuka
	logrus.Info("product service find by all start")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductsRepository.FindByAll(ctx, tx)

	//ini membuat logger penutup
	logrus.Info("product service find by all ended")

	return products
}
