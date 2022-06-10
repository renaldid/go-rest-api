package controller

import (
	"github.com/julienschmidt/httprouter"
	"go_rest_api/helper"
	"go_rest_api/model/web"
	"go_rest_api/service"
	"net/http"
	"strconv"
)

type CustomersControllerImpl struct {
	CustomersService service.CustomersService
}

func NewCustomersController(customersService service.CustomersService) CustomersController {
	return &CustomersControllerImpl{
		CustomersService: customersService,
	}
}

func (controller *CustomersControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customersCreateRequest := web.CustomersCreateRequest{}
	helper.ReadFromRequestBody(request, &customersCreateRequest)

	customersResponse := controller.CustomersService.Create(request.Context(), customersCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   customersResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomersControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customersUpdateRequest := web.CustomersUpdateRequest{}
	helper.ReadFromRequestBody(request, &customersUpdateRequest)

	customersId := params.ByName("customersId")
	id, err := strconv.Atoi(customersId)
	helper.PanicIfError(err)

	customersUpdateRequest.Id = id

	customersResponse := controller.CustomersService.Update(request.Context(), customersUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customersResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomersControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customersId := params.ByName("customersId")
	id, err := strconv.Atoi(customersId)
	helper.PanicIfError(err)

	controller.CustomersService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomersControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customersId := params.ByName("customersId")
	id, err := strconv.Atoi(customersId)
	helper.PanicIfError(err)

	customersResponse := controller.CustomersService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   customersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomersControllerImpl) FindByAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customersResponse := controller.CustomersService.FindByAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   customersResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
