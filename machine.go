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
