package helper

import (
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
)

func ToCustomersResponse(c domain.Customers) web.CustomersResponse {
	return web.CustomersResponse{
		Id:          c.Id,
		Name:        c.Name,
		Address:     c.Address,
		Email:       c.Email,
		PhoneNumber: c.PhoneNumber,
	}
}

func ToCustomersResponses(customer []domain.Customers) []web.CustomersResponse {
	var customersResponses []web.CustomersResponse
	for _, customers := range customer {
		customersResponses = append(customersResponses, ToCustomersResponse(customers))
	}
	return customersResponses
}
