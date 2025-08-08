package tcnsdk

import (
	"context"
	"fmt"
)

// RecognitionService handles visual recognition operations
type RecognitionService struct {
	client *Client
}

// ConfirmCommodity queries cloud product inventory
func (s *RecognitionService) ConfirmCommodity(ctx context.Context, req *CommodityConfirmRequest) (*CommodityConfirmResponse, error) {
	var resp CommodityConfirmResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/CommodityConfirm")

	if err != nil {
		return nil, fmt.Errorf("failed to confirm commodity: %w", err)
	}

	return &resp, nil
}

// RegisterProduct registers a new product for recognition
func (s *RecognitionService) RegisterProduct(ctx context.Context, req *ProductRegistrationRequest) (*ProductRegistrationResponse, error) {
	var resp ProductRegistrationResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Commodity/Apply")

	if err != nil {
		return nil, fmt.Errorf("failed to register product: %w", err)
	}

	return &resp, nil
}

// QueryProductReview queries the status of a product review
func (s *RecognitionService) QueryProductReview(ctx context.Context, req *ProductReviewQueryRequest) (*ProductReviewQueryResponse, error) {
	var resp ProductReviewQueryResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/Commodity/ApplyQuery")

	if err != nil {
		return nil, fmt.Errorf("failed to query product review: %w", err)
	}

	return &resp, nil
}

// SubmitRecognition submits a video for recognition
func (s *RecognitionService) SubmitRecognition(ctx context.Context, req *RecognitionRequest) (*RecognitionResponse, error) {
	var resp RecognitionResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/PushVideoAsk")

	if err != nil {
		return nil, fmt.Errorf("failed to submit recognition: %w", err)
	}

	return &resp, nil
}

// QueryRecognitionResult queries the result of a recognition task
func (s *RecognitionService) QueryRecognitionResult(
	ctx context.Context,
	req *RecognitionResultQueryRequest,
) (*RecognitionResultQueryResponse, error) {
	var resp RecognitionResultQueryResponse
	_, err := s.client.request(ctx).
		SetBody(req).
		SetResult(&resp).
		Post("/OpenApi/VideoAskResult")

	if err != nil {
		return nil, fmt.Errorf("failed to query recognition result: %w", err)
	}

	return &resp, nil
}

// GetRemainingQuota gets the remaining recognition quota
func (s *RecognitionService) GetRemainingQuota(ctx context.Context, appID string) (*RemainingQuotaResponse, error) {
	var resp RemainingQuotaResponse
	_, err := s.client.request(ctx).
		SetResult(&resp).
		Post(fmt.Sprintf("/OpenApi/GetVideoAskServiceNumber/%s", appID))

	if err != nil {
		return nil, fmt.Errorf("failed to get remaining quota: %w", err)
	}

	return &resp, nil
}
