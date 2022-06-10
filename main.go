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

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
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

	router := app.NewRouter(categoryController, customersController, ordersController, productsController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
