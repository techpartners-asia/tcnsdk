package tcnsdk

import (
	"context"
	"fmt"
)

type ProductService struct {
	client *Client
}

// NewProductService creates a new ProductService
func NewProductService(client *Client) *ProductService {
	return &ProductService{client: client}
}

// GetProduct returns a product by ID
func (s *ProductService) ListProducts(ctx context.Context, req *ProductListRequest) (*ProductListResponse, error) {
	var resp ProductListResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/CustomerCommoditys")

	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}

	return &resp, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *ProductUpdateRequest) (*ProductUpdateResponse, error) {
	var resp ProductUpdateResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/CustomerCommoditys/Update")

	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return &resp, nil
}
