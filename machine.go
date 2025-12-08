package tcnsdk

import (
	"context"
	"fmt"

	"github.com/techpartners-asia/tcnsdk/structs"
)

// MachineService handles machine-related operations
type MachineService struct {
	client *Client
}

// GetMachineInfo retrieves device information
func (s *MachineService) GetMachineInfo(ctx context.Context, machineID string) (*structs.MachineInfoResponse, error) {
	var resp structs.MachineInfoResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetHeader("Authorization", authResp.Data.Token).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/Machine/Info/%s", machineID))

	if err != nil {
		return nil, fmt.Errorf("failed to get machine info: %w", err)
	}

	return &resp, nil
}

// GetMachineDetail fetches machine information from /OpenApi/Machines/{mid}
func (s *MachineService) GetMachineDetail(ctx context.Context, machineID string) (*structs.MachineInfoResponse, error) {
	var resp structs.MachineInfoResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/Machines/%s", machineID))
	if err != nil {
		return nil, fmt.Errorf("failed to get machine detail: %w", err)
	}

	return &resp, nil
}

// ListDeviceTypes returns supported vending machine types
func (s *MachineService) ListDeviceTypes(ctx context.Context) (*structs.DeviceTypeResponse, error) {
	var resp structs.DeviceTypeResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetResult(&resp).
		Get("/OpenApi/VendTypes")
	if err != nil {
		return nil, fmt.Errorf("failed to list device types: %w", err)
	}

	return &resp, nil
}

// ListMachines paginates machines using /OpenApi/Machines
func (s *MachineService) ListMachines(ctx context.Context, req *structs.MachineListRequest) (*structs.MachineListResponse, error) {
	var resp structs.MachineListResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	params := map[string]string{
		"pageIndex": fmt.Sprintf("%d", req.PageIndex),
		"pageSize":  fmt.Sprintf("%d", req.PageSize),
	}
	if req.MachineNoOrName != "" {
		params["machineNoOrName"] = req.MachineNoOrName
	}
	if req.OnlineStatus != nil {
		params["onlineStatus"] = fmt.Sprintf("%d", *req.OnlineStatus)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetQueryParams(params).
		SetResult(&resp).
		Get("/OpenApi/Machines")
	if err != nil {
		return nil, fmt.Errorf("failed to list machines: %w", err)
	}

	return &resp, nil
}

// SetMachineAttributes updates machine temperature/volume attributes
func (s *MachineService) SetMachineAttributes(ctx context.Context, machineID string, req *structs.MachineControlRequest) (*structs.MachineControlResponse, error) {
	var resp structs.MachineControlResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post(fmt.Sprintf("/OpenApi/Machines/Attributes/%s", machineID))
	if err != nil {
		return nil, fmt.Errorf("failed to set machine attributes: %w", err)
	}

	return &resp, nil
}

// GetMachineCommodities retrieves equipment product information
func (s *MachineService) GetMachineCommodities(ctx context.Context, machineID string) (*structs.CommodityResponse, error) {
	var resp structs.CommodityResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/Machine/Commoditys/%s", machineID))

	if err != nil {
		return nil, fmt.Errorf("failed to get machine commodities: %w", err)
	}

	return &resp, nil
}

// GetMachinePayConfig fetches payment configuration for a device
func (s *MachineService) GetMachinePayConfig(ctx context.Context, machineID string) (*structs.PayConfigResponse, error) {
	var resp structs.PayConfigResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetResult(&resp).
		Get(fmt.Sprintf("/OpenApi/Machine/PayConfig/%s", machineID))
	if err != nil {
		return nil, fmt.Errorf("failed to get machine pay config: %w", err)
	}

	return &resp, nil
}

// ListSlot lists the slots in a vending machine with their product information
func (s *MachineService) ListSlot(ctx context.Context, vendId string) (*structs.ListSlotResponse, error) {
	var resp structs.ListSlotResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
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
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
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
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
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
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/VendSlotCommoditys/Modify")
	if err != nil {
		return nil, fmt.Errorf("failed to update product on machine: %w", err)
	}

	return &resp, nil
}
