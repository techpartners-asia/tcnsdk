package tcnsdk

import (
	"context"
	"fmt"

	"github.com/techpartners-asia/tcnsdk/structs"
)

// MachineService handles machine-related operations
type MachineService struct {
	Client *Client
}

// GetMachineInfo retrieves device information
func (s *MachineService) GetMachineInfo(ctx context.Context, machineID string) (*structs.MachineInfoResponse, error) {
	var resp structs.MachineInfoResponse
	_, err := s.Client.Request(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/Machine/Info/%s", machineID))

	if err != nil {
		return nil, fmt.Errorf("failed to get machine info: %w", err)
	}

	return &resp, nil
}

// GetMachineCommodities retrieves equipment product information
func (s *MachineService) GetMachineCommodities(ctx context.Context, machineID string) (*structs.CommodityResponse, error) {
	var resp structs.CommodityResponse
	_, err := s.Client.Request(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/Machine/Commoditys/%s", machineID))

	if err != nil {
		return nil, fmt.Errorf("failed to get machine commodities: %w", err)
	}

	return &resp, nil
}

// ListSlot lists the slots in a vending machine with their product information
func (s *MachineService) ListSlot(ctx context.Context, vendId string) (*structs.ListSlotResponse, error) {
	var resp structs.ListSlotResponse
	_, err := s.Client.Request(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/%s/VendSlotCommoditys", vendId))
	if err != nil {
		return nil, fmt.Errorf("failed to list slot: %w", err)
	}
	return &resp, nil
}

// AddProductToMachine adds a product to a vending machine
func (s *MachineService) AddProductToMachine(ctx context.Context, req *structs.AddProductToMachineRequest) (*structs.AddProductToMachineResponse, error) {
	var resp structs.AddProductToMachineResponse
	_, err := s.Client.Request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/VendSlotCommoditys/Add")

	if err != nil {
		return nil, fmt.Errorf("failed to add product to machine: %w", err)
	}

	return &resp, nil
}

// DeleteProductFromMachine removes a product from a vending machine
func (s *MachineService) DeleteProductFromMachine(ctx context.Context, machineID string, req *structs.DeleteProductFromMachineRequest) (*structs.DeleteProductFromMachineResponse, error) {
	var resp structs.DeleteProductFromMachineResponse
	_, err := s.Client.Request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/VendSlotCommoditys/Delete")
	if err != nil {
		return nil, fmt.Errorf("failed to delete product from machine: %w", err)
	}

	return &resp, nil
}

// UpdateProductOnMachine updates a product on a vending machine
func (s *MachineService) UpdateProductOnMachine(ctx context.Context, machineID string, req *structs.UpdateProductOnMachineRequest) (*structs.UpdateProductOnMachineResponse, error) {
	var resp structs.UpdateProductOnMachineResponse
	_, err := s.Client.Request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/VendSlotCommoditys/Modify")
	if err != nil {
		return nil, fmt.Errorf("failed to update product on machine: %w", err)
	}

	return &resp, nil
}
