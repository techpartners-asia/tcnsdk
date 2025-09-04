package tcnsdk

import (
	"context"
	"fmt"
)

type TrainService struct {
	client *Client
}

func NewTrainService(client *Client) *TrainService {
	return &TrainService{client: client}
}

func (s *TrainService) TrainProduct(ctx context.Context, req *ProductTrainRequest) (*ProductTrainResponse, error) {
	var resp ProductTrainResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/PutNewCommodity")

	if err != nil {
		return nil, fmt.Errorf("failed to train product: %w", err)
	}

	return &resp, nil
}

func (s *TrainService) ListCheckProductTrain(ctx context.Context, req *ListCheckProductTrainRequest) (*ListCheckProductTrainResponse, error) {
	var resp ListCheckProductTrainResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/CommodityApply/Record")

	if err != nil {
		return nil, fmt.Errorf("failed to list check product train: %w", err)
	}

	return &resp, nil
}
