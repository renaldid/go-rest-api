package service

import (
	"context"
	"go_rest_api/model/web"
)

type ProductsService interface {
	Create(ctx context.Context, request web.ProductsCreateRequest) web.ProductsResponse
	Update(ctx context.Context, request web.ProductsUpdateRequest) web.ProductsResponse
	Delete(ctx context.Context, productsId int)
	FindById(ctx context.Context, productsId int) web.ProductsResponse
	FindByAll(ctx context.Context) []web.ProductsResponse
}
