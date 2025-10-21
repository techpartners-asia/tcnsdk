package tcnsdk

import (
	"fmt"

	"github.com/techpartners-asia/tcnsdk/structs"
)

type TrainService struct {
	client *Client
}

func NewTrainService(client *Client) *TrainService {
	return &TrainService{client: client}
}

func (s *TrainService) TrainProduct(req *structs.ProductTrainRequest) (*structs.ProductTrainResponse, error) {
	var resp structs.ProductTrainResponse
	_, err := s.client.Request().
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/PutNewCommodity")

	if err != nil {
		return nil, fmt.Errorf("failed to train product: %w", err)
	}

	return &resp, nil
}

func (s *TrainService) ListProductTrainRequest(req *structs.ListProductTrainRequest) (*structs.ListProductTrainResponse, error) {
	var resp structs.ListProductTrainResponse
	_, err := s.client.Request().
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/CommodityApply/Record")

	if err != nil {
		return nil, fmt.Errorf("failed to list product train request: %w", err)
	}

	return &resp, nil
}
