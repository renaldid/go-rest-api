package service

import (
	"context"
	"go_rest_api/model/web"
)

type OrderProductsService interface {
	Create(ctx context.Context, request web.OrderProductCreateRequest) web.OrderProductResponse
	Update(ctx context.Context, request web.OrderProductUpdateRequest) web.OrderProductResponse
	Delete(ctx context.Context, orderProductId int)
	FindById(ctx context.Context, orderProductId int) web.OrderProductResponse
	FindByAll(ctx context.Context) []web.OrderProductResponse
}
