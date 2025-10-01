package tcnsdk

import (
	"context"
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
func (s *OrderService) OpenDoor(ctx context.Context, req *structs.OpenDoorRequest) (*structs.OpenDoorResponse, error) {
	if req.TimeSp == 0 {
		req.TimeSp = time.Now().Unix()
	}

	var resp structs.OpenDoorResponse
	_, err := s.Client.Request(ctx).
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

	var resp structs.RestockOpenDoorResponse
	_, err := s.Client.Request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Repli/OpenDoorMethod")

	if err != nil {
		return nil, fmt.Errorf("failed to open restock door: %w", err)
	}

	return &resp, nil
}
