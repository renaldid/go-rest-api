package controller

import (
	"github.com/julienschmidt/httprouter"
	"go_rest_api/helper"
	"go_rest_api/model/web"
	"go_rest_api/service"
	"net/http"
	"strconv"
)

type OrderProductsControllerImpl struct {
	ProductsService service.OrderProductsService
}

func NewOrderProductsController(ProductsService service.OrderProductsService) OrderProductsController {
	return &OrderProductsControllerImpl{
		ProductsService: ProductsService,
	}
}

func (controller *OrderProductsControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	productsCreateRequest := web.OrderProductCreateRequest{}
	helper.ReadFromRequestBody(request, &productsCreateRequest)

	productsResponse := controller.ProductsService.Create(request.Context(), productsCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   productsResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductsControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderProductsUpdateRequest := web.OrderProductUpdateRequest{}
	helper.ReadFromRequestBody(request, &orderProductsUpdateRequest)

	productsId := params.ByName("productsId")
	id, err := strconv.Atoi(productsId)
	helper.PanicIfError(err)

	orderProductsUpdateRequest.Id = id
	ordersResponse := controller.ProductsService.Update(request.Context(), orderProductsUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ordersResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductsControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productsId := params.ByName("productsId")
	id, err := strconv.Atoi(productsId)
	helper.PanicIfError(err)

	controller.ProductsService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderProductsControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
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
}

func (controller *OrderProductsControllerImpl) FindByAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productsResponse := controller.ProductsService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   productsResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)

}
