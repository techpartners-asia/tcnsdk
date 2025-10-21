package tcnsdk

import (
	"fmt"

	"github.com/techpartners-asia/tcnsdk/structs"
)

type ProductService struct {
	client *Client
}

// NewProductService creates a new ProductService
func NewProductService(client *Client) *ProductService {
	return &ProductService{client: client}
}

// GetProduct returns a product by ID
func (s *ProductService) ListProducts(req *structs.ProductListRequest) (*structs.ProductListResponse, error) {
	var resp structs.ProductListResponse
	authResp, err := s.client.getAuthResponse()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetAuthToken(authResp.Data.Token).
		SetQueryParams(map[string]string{
			"pageIndex":     fmt.Sprintf("%d", req.PageIndex),
			"pageSize":      fmt.Sprintf("%d", req.PageSize),
			"name":          req.Name,
			"commodityType": req.CommodityType,
			"supportType":   req.SupportType,
		}).
		SetResult(&resp).
		Get("/OpenApi/CustomerCommoditys")

	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}

	return &resp, nil
}

func (s *ProductService) UpdateProduct(req *structs.ProductUpdateRequest) (*structs.ProductUpdateResponse, error) {
	var resp structs.ProductUpdateResponse
	authResp, err := s.client.getAuthResponse()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/CustomerCommoditys/Update")

	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return &resp, nil
}
