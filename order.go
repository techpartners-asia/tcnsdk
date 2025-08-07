package tcnsdk

import (
	"context"
	"fmt"
	"time"
)

// OrderService handles order-related operations
type OrderService struct {
	client *Client
}

// OpenDoorRequest represents door opening request
type OpenDoorRequest struct {
	OrderID         string `json:"orderId"`
	MachineID       string `json:"machineId"`
	DoorNo          int    `json:"doorNo"`
	TranseType      int    `json:"transeType"`
	CustomerDetails string `json:"customerDetails,omitempty"`
	TimeSp          int64  `json:"timeSp"`
	NotifyURL       string `json:"NotifyUrl"`
	Remark          string `json:"remark,omitempty"`
}

// OpenDoorResponse represents door opening response
type OpenDoorResponse struct {
	OrderID    string `json:"OrderId"`
	TranseType int    `json:"TranseType"`
	Action     string `json:"Action"`
	Status     bool   `json:"Status"`
	Msg        string `json:"Msg"`
	Data       string `json:"Data"`
}

// RestockOpenDoorRequest represents restock door opening request
type RestockOpenDoorRequest struct {
	OrderID         string `json:"orderId"`
	MachineID       string `json:"machineId"`
	DoorNo          int    `json:"doorNo"`
	TranseType      int    `json:"transeType"`
	CustomerDetails string `json:"customerDetails,omitempty"`
	TimeSp          int64  `json:"timeSp"`
	NotifyURL       string `json:"NotifyUrl"`
	Remark          string `json:"remark,omitempty"`
}

// RestockOpenDoorResponse represents restock door opening response
type RestockOpenDoorResponse struct {
	OrderID    string `json:"OrderId"`
	TranseType int    `json:"TranseType"`
	Action     string `json:"Action"`
	Status     bool   `json:"Status"`
	Msg        string `json:"Msg"`
	Data       string `json:"Data"`
}

// OpenDoor opens the vending machine door
func (s *OrderService) OpenDoor(ctx context.Context, req *OpenDoorRequest) (*OpenDoorResponse, error) {
	if req.TimeSp == 0 {
		req.TimeSp = time.Now().Unix()
	}

	var resp OpenDoorResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Order/OpenDoor")

	if err != nil {
		return nil, fmt.Errorf("failed to open door: %w", err)
	}

	return &resp, nil
}

// RestockOpenDoor opens the door for restocking (testing endpoint)
func (s *OrderService) RestockOpenDoor(ctx context.Context, req *RestockOpenDoorRequest) (*RestockOpenDoorResponse, error) {
	if req.TimeSp == 0 {
		req.TimeSp = time.Now().Unix()
	}

	var resp RestockOpenDoorResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Repli/OpenDoorMethod")

	if err != nil {
		return nil, fmt.Errorf("failed to open restock door: %w", err)
	}

	return &resp, nil
}

// ReportPaymentResult reports the payment result
func (s *OrderService) ReportPaymentResult(ctx context.Context, req *ReportPaymentResultRequest) (*ReportPaymentResultResponse, error) {
	var resp ReportPaymentResultResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Order/ReportPaymentResult")

	if err != nil {
		return nil, fmt.Errorf("failed to report payment result: %w", err)
	}

	return &resp, nil
}

// ConfirmReplenishment confirms replenishment data
func (s *OrderService) ConfirmReplenishment(ctx context.Context, req *ConfirmRepliRequest) (*ConfirmRepliResponse, error) {
	var resp ConfirmRepliResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Repli/Confirm")

	if err != nil {
		return nil, fmt.Errorf("failed to confirm replenishment: %w", err)
	}

	return &resp, nil
}
