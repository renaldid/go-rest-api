package helper

import (
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
)

func ToOrderProductResponse(o domain.OrdersProduct) web.OrderProductResponse {
	return web.OrderProductResponse{
		Id:         o.Id,
		OrderId:    o.OrderId,
		ProductId:  o.ProductId,
		Qty:        o.Qty,
		Amount:     o.Amount,
		CreatedAt:  o.CreatedAt,
		UploadedAt: o.UploadedAt,
	}
}

func ToOrderProduct(o web.OrderProductResponse) domain.OrdersProduct {
	return domain.OrdersProduct{
		Id:         o.Id,
		OrderId:    o.OrderId,
		ProductId:  o.ProductId,
		Qty:        o.Qty,
		Amount:     o.Amount,
		CreatedAt:  o.CreatedAt,
		UploadedAt: o.UploadedAt,
	}
}

func ToOrderProductResponses(o []domain.OrdersProduct) []web.OrderProductResponse {
	var orderProductResponses []web.OrderProductResponse
	for _, orderProduct := range o {
		orderProductResponses = append(orderProductResponses, ToOrderProductResponse(orderProduct))
	}
	return orderProductResponses
}
