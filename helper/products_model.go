package helper

import (
	"go_rest_api/model/domain"
	"go_rest_api/model/web"
)

func ToProductsResponse(p domain.Products) web.ProductsResponse {
	return web.ProductsResponse{
		Id:         p.Id,
		Name:       p.Name,
		Price:      p.Price,
		CategoryId: p.CategoryId,
	}
}
func ToProducts(p web.ProductsResponse) domain.Products {
	return domain.Products{
		Id:         p.Id,
		Name:       p.Name,
		Price:      p.Price,
		CategoryId: p.CategoryId,
	}
}

func ToProductsResponses(products []domain.Products) []web.ProductsResponse {
	var productResponses []web.ProductsResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductsResponse(product))
	}
	return productResponses
}
