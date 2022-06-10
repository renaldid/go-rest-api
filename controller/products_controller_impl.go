package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"go_rest_api/helper"
	"go_rest_api/model/web"
	"go_rest_api/service"
	"net/http"
	"strconv"
)

type ProductsControllerImpl struct {
	ProductsService service.ProductsService
}

func NewProductsController(productsService service.ProductsService) ProductsController {
	return &ProductsControllerImpl{
		ProductsService: productsService,
	}
}

func (controller *ProductsControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//ini membuat logger pembuka
	logrus.Info("product controller create start")

	productsCreateRequest := web.ProductsCreateRequest{}
	helper.ReadFromRequestBody(request, &productsCreateRequest)

	productsResponse := controller.ProductsService.Create(request.Context(), productsCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   productsResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
	//ini membuat logger penutup
	logrus.Info("product controller create ended")
}

func (controller *ProductsControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//ini membuat logger pembuka
	logrus.Info("product controller update start")

	productsUpdateRequest := web.ProductsUpdateRequest{}
	helper.ReadFromRequestBody(request, &productsUpdateRequest)

	productsId := params.ByName("productsId")
	id, err := strconv.Atoi(productsId)
	helper.PanicIfError(err)

	productsUpdateRequest.Id = id
	ordersResponse := controller.ProductsService.Update(request.Context(), productsUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ordersResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
	//ini membuat logger penutup
	logrus.Info("product controller update ended")
}

func (controller *ProductsControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//ini membuat logger pembuka
	logrus.Info("product controller delete start")

	productsId := params.ByName("productsId")
	id, err := strconv.Atoi(productsId)
	helper.PanicIfError(err)

	controller.ProductsService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
	//ini membuat logger penutup
	logrus.Info("product controller delete ended")
}

func (controller *ProductsControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//ini membuat logger pembuka
	logrus.Info("product controller find by id start")

	productsId := params.ByName("productsId")
	id, err := strconv.Atoi(productsId)
	helper.PanicIfError(err)

	productsResponse := controller.ProductsService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   productsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
	//ini membuat logger penutup
	logrus.Info("product controller find by id ended")
}

func (controller *ProductsControllerImpl) FindByAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//ini membuat logger pembuka
	logrus.Info("product controller find by all start")

	productsResponse := controller.ProductsService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   productsResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
	//ini membuat logger penutup
	logrus.Info("product controller find by all ended")
}
