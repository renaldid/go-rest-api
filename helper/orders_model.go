package helper

import (
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
)

func ToOrdersResponse(o domain.Orders) web.OrdersResponse {
	return web.OrdersResponse{
		Id:          o.Id,
		CustomerId:  o.CustomerId,
		TotalAmount: o.TotalAmount,
	}
}

func ToOrdersResponses(order []domain.Orders) []web.OrdersResponse {
	var ordersResponses []web.OrdersResponse
	for _, orders := range order {
		ordersResponses = append(ordersResponses, ToOrdersResponse(orders))
	}
	return ordersResponses
}
