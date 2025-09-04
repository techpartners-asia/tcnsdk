package tcnsdk

import (
	"context"
	"fmt"
)

// MachineService handles machine-related operations
type MachineService struct {
	client *Client
}

// GetMachineInfo retrieves device information
func (s *MachineService) GetMachineInfo(ctx context.Context, machineID string) (*MachineInfoResponse, error) {
	var resp MachineInfoResponse
	_, err := s.client.request(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/Machine/Info/%s", machineID))

	if err != nil {
		return nil, fmt.Errorf("failed to get machine info: %w", err)
	}

	return &resp, nil
}

// GetMachineCommodities retrieves equipment product information
func (s *MachineService) GetMachineCommodities(ctx context.Context, machineID string) (*CommodityResponse, error) {
	var resp CommodityResponse
	_, err := s.client.request(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/Machine/Commoditys/%s", machineID))

	if err != nil {
		return nil, fmt.Errorf("failed to get machine commodities: %w", err)
	}

	return &resp, nil
}

func (s *MachineService) ListSlot(ctx context.Context, vendId string) (*ListSlotResponse, error) {
	var resp ListSlotResponse
	_, err := s.client.request(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/%s/VendSlotCommoditys", vendId))
	if err != nil {
		return nil, fmt.Errorf("failed to list slot: %w", err)
	}
	return &resp, nil
}

func (s *MachineService) AddProductToMachine(ctx context.Context, req *AddProductToMachineRequest) (*AddProductToMachineResponse, error) {
	var resp AddProductToMachineResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/VendSlotCommoditys/Add")
	if err != nil {
		return nil, fmt.Errorf("failed to add product to machine: %w", err)
	}

	return &resp, nil
}

func (s *MachineService) DeleteProductFromMachine(ctx context.Context, machineID string, req *DeleteProductFromMachineRequest) (*DeleteProductFromMachineResponse, error) {
	var resp DeleteProductFromMachineResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/VendSlotCommoditys/Delete")
	if err != nil {
		return nil, fmt.Errorf("failed to delete product from machine: %w", err)
	}

	return &resp, nil
}

func (s *MachineService) UpdateProductOnMachine(ctx context.Context, machineID string, req *UpdateProductOnMachineRequest) (*UpdateProductOnMachineResponse, error) {
	var resp UpdateProductOnMachineResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/VendSlotCommoditys/Modify")
	if err != nil {
		return nil, fmt.Errorf("failed to update product on machine: %w", err)
	}

	return &resp, nil
}

// GetMachinePayConfig retrieves device payment configuration information
func (s *MachineService) GetMachinePayConfig(ctx context.Context, machineID string) (*PayConfigResponse, error) {
	var resp PayConfigResponse
	_, err := s.client.request(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/Machine/PayConfig/%s", machineID))

	if err != nil {
		return nil, fmt.Errorf("failed to get machine pay config: %w", err)
	}

	return &resp, nil
}
