package tcnsdk

import (
	"fmt"
	"time"

	"github.com/techpartners-asia/tcnsdk/structs"
)

// OrderService handles order-related operations
type OrderService struct {
	Client *Client
}

// OpenDoor opens the vending machine door
// [stable] [tested]
func (s *OrderService) OpenDoor(req *structs.OpenDoorRequest) (*structs.OpenDoorResponse, error) {
	if req.TimeSp == 0 {
		req.TimeSp = time.Now().Unix()
	}

	authResp, err := s.Client.getAuthResponse()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	var resp structs.OpenDoorResponse
	_, err = s.Client.client.R().
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
func (s *OrderService) RestockOpenDoor(req *structs.RestockOpenDoorRequest) (*structs.RestockOpenDoorResponse, error) {
	if req.TimeSp == 0 {
		req.TimeSp = time.Now().Unix()
	}

	authResp, err := s.Client.getAuthResponse()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	var resp structs.RestockOpenDoorResponse
	_, err = s.Client.client.R().
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Repli/OpenDoorMethod")

	if err != nil {
		return nil, fmt.Errorf("failed to open restock door: %w", err)
	}

	return &resp, nil
}

// RestockOpenDoor opens the door for restocking (testing endpoint)
func (s *OrderService) OrderDetail(transID string) (*structs.OrderDetailResponse, error) {
	var resp structs.OrderDetailResponse
	authResp, err := s.Client.getAuthResponse()
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.Client.client.R().
		SetAuthToken(authResp.Data.Token).
		SetPathParam("transID", transID).
		SetResult(&resp).
		Get("/OpenApi/Order/{transID}/Detail")

	if err != nil {
		return nil, fmt.Errorf("failed to get order detail: %w", err)
	}

	return &resp, nil
}
