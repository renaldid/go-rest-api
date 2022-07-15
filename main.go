package main

import (
	"github.com/go-playground/validator"
	"go_rest_api/app"
	"go_rest_api/controller"
	"go_rest_api/helper"
	"go_rest_api/middleware"
	"go_rest_api/repository"
	"go_rest_api/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	//1. Membuat objek bernama categoryRepository
	//2. Memanggil constructor bernama NewCategoryRepository
	//3. Tanpa ada dependency
	categoryRepository := repository.NewCategoryRepository()
	//1. Membuat object bernama categoryService
	//2. Memanggil constructor bernama NewCategoryService
	//3. Meng-inject dependency bernama categoryRepository, db, validate
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//1. Membuat object bernama categoryController
	//2. Memanggil constructor bernama NewCategoryController
	//3. Meng-inject dependency bernama categoryService
	categoryController := controller.NewCategoryController(categoryService)

	customersRepo := repository.NewCustomersRepository()
	customersService := service.NewCustomersService(customersRepo, db, validate)
	customersController := controller.NewCustomersController(customersService)

	ordersRepo := repository.NewOrdersRepository()
	ordersService := service.NewOrdersService(ordersRepo, db, validate)
	ordersController := controller.NewOrderController(ordersService)

	productsRepo := repository.NewProductsRepository()
	productsService := service.NewProductsService(productsRepo, db, validate)
	productsController := controller.NewProductsController(productsService)

	orderProductsRepo := repository.NewOrderProductRepository()
	orderProductService := service.NewOrderProductsService(orderProductsRepo, db, validate)
	orderProductController := controller.NewOrderProductsController(orderProductService)
	router := app.NewRouter(categoryController, customersController, ordersController, productsController, orderProductController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
