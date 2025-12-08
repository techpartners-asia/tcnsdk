package tcnsdk

import (
	"context"
	"fmt"

	"github.com/techpartners-asia/tcnsdk/structs"
)

type TrainService struct {
	client *Client
}

func NewTrainService(client *Client) *TrainService {
	return &TrainService{client: client}
}

func (s *TrainService) TrainProduct(ctx context.Context, req *structs.ProductTrainRequest) (*structs.ProductTrainResponse, error) {
	var resp structs.ProductTrainResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/PutNewCommodity")

	if err != nil {
		return nil, fmt.Errorf("failed to train product: %w", err)
	}

	return &resp, nil
}

func (s *TrainService) ListProductTrainRequest(ctx context.Context, req *structs.ListProductTrainRequest) (*structs.ListProductTrainResponse, error) {
	var resp structs.ListProductTrainResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetBody(req).
		SetQueryParam("recordId", fmt.Sprintf("%d", req.RecordID)).
		SetResult(&resp).
		Post("/OpenApi/CommodityApply/Record")

	if err != nil {
		return nil, fmt.Errorf("failed to list product train request: %w", err)
	}

	return &resp, nil
}
func (s *TrainService) GetTrainRequest(ctx context.Context, id string) (*structs.ProductTrain, error) {
	var resp structs.ListProductTrainResponse
	authResp, err := s.client.getAuthResponse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get auth response: %w", err)
	}

	_, err = s.client.client.R().
		SetContext(ctx).
		SetAuthToken(authResp.Data.Token).
		SetQueryParam("recordId", id).
		SetResult(&resp).
		Get("/OpenApi/CommodityApply/Record")

	if err != nil {
		return nil, fmt.Errorf("failed to list product train request: %w", err)
	}

	if len(resp.Data.Items) == 0 {
		return nil, structs.ErrNoProductTrainRequestFound
	}

	item := resp.Data.Items[0]

	return &item, nil
}
