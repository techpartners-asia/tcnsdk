package tcnsdk

import (
	"context"
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

// GetProduct returns a product by ID or SKU
// you can pass either productID or sku
func (s *ProductService) GetProduct(ctx context.Context, productIDorSku string) (*structs.Product, error) {
	var resp structs.Product
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetPathParam("productIDorSku", productIDorSku).
		SetResult(&resp).
		Get("/OpenApi/CustomerCommoditys/{productIDorSku}")

	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return &resp, nil
}

// ListProducts returns a list of products
func (s *ProductService) ListProducts(ctx context.Context, req *structs.ProductListRequest) (*structs.ProductListResponse, error) {
	var resp structs.ProductListResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
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

// UpdateProduct updates product information
func (s *ProductService) UpdateProduct(ctx context.Context, req *structs.ProductUpdateRequest) (*structs.ProductUpdateResponse, error) {
	var resp structs.ProductUpdateResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/CustomerCommoditys/Update")

	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return &resp, nil
}

// UploadProduct registers a new product (commodity) in the catalog
func (s *ProductService) UploadProduct(ctx context.Context, req *structs.ProductUploadRequest) (*structs.ProductUploadResponse, error) {
	var resp structs.ProductUploadResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/PutNewCommodity")
	if err != nil {
		return nil, fmt.Errorf("failed to upload product: %w", err)
	}

	return &resp, nil
}

// ListProductApplications lists product onboarding applications (CommodityApplys)
func (s *ProductService) ListProductApplications(ctx context.Context, req *structs.ProductApplyListRequest) (*structs.ProductApplyListResponse, error) {
	var resp structs.ProductApplyListResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	params := map[string]string{
		"pageIndex": fmt.Sprintf("%d", req.PageIndex),
		"pageSize":  fmt.Sprintf("%d", req.PageSize),
	}
	if req.Sku != "" {
		params["sku"] = req.Sku
	}
	if req.State != nil {
		params["state"] = fmt.Sprintf("%d", *req.State)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetQueryParams(params).
		SetResult(&resp).
		Get("/OpenApi/CommodityApplys")
	if err != nil {
		return nil, fmt.Errorf("failed to list product applications: %w", err)
	}

	return &resp, nil
}

// GetProductApplicationDetail fetches a single CommodityApply by SKU/ID
func (s *ProductService) GetProductApplicationDetail(ctx context.Context, sku string) (*structs.ProductApplyDetailResponse, error) {
	var resp structs.ProductApplyDetailResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetPathParam("sku", sku).
		SetResult(&resp).
		Get("/OpenApi/CommodityApplys/{sku}")
	if err != nil {
		return nil, fmt.Errorf("failed to get product application detail: %w", err)
	}

	return &resp, nil
}
