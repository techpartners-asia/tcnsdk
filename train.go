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

func (s *TrainService) ListProductTrainRequest(ctx context.Context, req *ListProductTrainRequest) (*ListProductTrainResponse, error) {
	var resp ListProductTrainResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/CommodityApply/Record")

	if err != nil {
		return nil, fmt.Errorf("failed to list product train request: %w", err)
	}

	return &resp, nil
}
