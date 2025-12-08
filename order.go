package tcnsdk

import (
	"context"
	"fmt"
	"time"

	"github.com/techpartners-asia/tcnsdk/structs"
)

// OrderService handles order-related operations
type OrderService struct {
	client *Client
}

// OpenDoor opens the vending machine door
// [stable] [tested]
func (s *OrderService) OpenDoor(ctx context.Context, req *structs.OpenDoorRequest) (*structs.OpenDoorResponse, error) {
	if req.TimeSp == 0 {
		req.TimeSp = time.Now().Unix()
	}

	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	var resp structs.OpenDoorResponse
	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Order/OpenDoor")

	if err != nil {
		return nil, fmt.Errorf("failed to open door: %w", err)
	}

	return &resp, nil
}

// RestockOpenDoor opens the door for restocking (testing endpoint)
func (s *OrderService) RestockOpenDoor(ctx context.Context, req *structs.RestockOpenDoorRequest) (*structs.RestockOpenDoorResponse, error) {
	if req.TimeSp == 0 {
		req.TimeSp = time.Now().Unix()
	}

	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	var resp structs.RestockOpenDoorResponse
	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Repli/OpenDoorMethod")

	if err != nil {
		return nil, fmt.Errorf("failed to open restock door: %w", err)
	}

	return &resp, nil
}

// OrderDetail retrieves order detail by transaction ID
func (s *OrderService) OrderDetail(ctx context.Context, transID string) (*structs.OrderDetailResponse, error) {
	var resp structs.OrderDetailResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetPathParam("transID", transID).
		SetResult(&resp).
		Get("/OpenApi/Order/{transID}/Detail")

	if err != nil {
		return nil, fmt.Errorf("failed to get order detail: %w", err)
	}

	return &resp, nil
}

// GetOpenDoorStatus polls the status for a door-open transaction
func (s *OrderService) GetOpenDoorStatus(ctx context.Context, transID string) (*structs.OpenDoorStatusResponse, error) {
	var resp structs.OpenDoorStatusResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetPathParam("transId", transID).
		SetResult(&resp).
		Get("/OpenApi/Order/OpenDoor/{transId}")
	if err != nil {
		return nil, fmt.Errorf("failed to get open door status: %w", err)
	}

	return &resp, nil
}

// GetOrderVideo retrieves the video evidence for a transaction
func (s *OrderService) GetOrderVideo(ctx context.Context, transID string) (*structs.OrderVideoResponse, error) {
	var resp structs.OrderVideoResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetPathParam("transId", transID).
		SetResult(&resp).
		Get("/OpenApi/Order/Video/{transId}")
	if err != nil {
		return nil, fmt.Errorf("failed to get order video: %w", err)
	}

	return &resp, nil
}

// ConfirmReplenishment submits replenishment confirmation data
func (s *OrderService) ConfirmReplenishment(ctx context.Context, req *structs.ConfirmRepliRequest) (*structs.ConfirmRepliResponse, error) {
	var resp structs.ConfirmRepliResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Repli/Confirm")
	if err != nil {
		return nil, fmt.Errorf("failed to confirm replenishment: %w", err)
	}

	return &resp, nil
}

// ReportPaymentResult notifies SaaS about order payment status
func (s *OrderService) ReportPaymentResult(ctx context.Context, req *structs.ReportPaymentResultRequest) (*structs.ReportPaymentResultResponse, error) {
	var resp structs.ReportPaymentResultResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Order/ReportPaymentResult")
	if err != nil {
		return nil, fmt.Errorf("failed to report payment result: %w", err)
	}

	return &resp, nil
}

// ReportPaymentResultV2 sends detailed payment results with product breakdowns
func (s *OrderService) ReportPaymentResultV2(ctx context.Context, req *structs.ReportPaymentResultV2Request) (*structs.ReportPaymentResultV2Response, error) {
	var resp structs.ReportPaymentResultV2Response
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Order/V2/ReportPaymentResult")
	if err != nil {
		return nil, fmt.Errorf("failed to report payment result v2: %w", err)
	}

	return &resp, nil
}
