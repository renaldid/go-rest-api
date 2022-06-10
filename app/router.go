package app

import (
	"github.com/julienschmidt/httprouter"
	"go_rest_api/controller"
	"go_rest_api/exception"
)

func NewRouter(categoryController controller.CategoryController, customersController controller.CustomersController, ordersController controller.OrdersController, productsController controller.ProductsController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindByAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories/", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/customers", customersController.FindByAll)
	router.GET("/api/customers/:customersId", customersController.FindById)
	router.POST("/api/customers/", customersController.Create)
	router.PUT("/api/customers/:customersId", customersController.Update)
	router.DELETE("/api/customers/:customersId", customersController.Delete)

	router.GET("/api/orders", ordersController.FindByAll)
	router.GET("/api/orders/:ordersId", ordersController.FindById)
	router.POST("/api/orders/", ordersController.Create)
	router.PUT("/api/orders/:ordersId", ordersController.Update)
	router.DELETE("/api/orders/:ordersId", ordersController.Delete)

	router.GET("/api/products", productsController.FindByAll)
	router.GET("/api/products/:productsId", productsController.FindById)
	router.POST("/api/products/", productsController.Create)
	router.PUT("/api/products/:productsId", productsController.Update)
	router.DELETE("/api/products/:ordersId", productsController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
