package tcnsdk

import (
	"context"
	"time"

	"resty.dev/v3"
)

// Client represents the SAAS API client
type Client struct {
	client *resty.Client
	config *Config

	// Services
	Machine     *MachineService
	Order       *OrderService
	Recognition *RecognitionService
}

// Config holds the client configuration
type Config struct {
	BaseURL string
	AppID   string
	Key     string
	Secret  string
	Timeout time.Duration
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseURL: "https://openapi1.ourvend.com",
		Timeout: 30 * time.Second,
	}
}

// NewClient creates a new SAAS API client
func NewClient(config *Config) *Client {
	if config == nil {
		config = DefaultConfig()
	}

	client := resty.New().
		SetBaseURL(config.BaseURL).
		SetTimeout(config.Timeout)

	c := &Client{
		client: client,
		config: config,
	}

	// Initialize services
	c.Machine = &MachineService{client: c}
	c.Order = &OrderService{client: c}
	c.Recognition = &RecognitionService{client: c}

	return c
}

// getAuthToken retrieves authentication token
func (c *Client) getAuthToken(ctx context.Context) (string, error) {
	authReq := &AuthRequest{
		AppID:  c.config.AppID,
		Key:    c.config.Key,
		Secret: c.config.Secret,
	}

	var authResp AuthResponse
	_, err := c.client.R().
		SetContext(ctx).
		SetBody(authReq).
		SetResult(&authResp).
		Post("/OpenApi/Login")

	if err != nil {
		return "", err
	}

	return authResp.Token, nil
}

// request performs an authenticated request
func (c *Client) request(ctx context.Context) *resty.Request {
	token, err := c.getAuthToken(ctx)
	if err != nil {
		// Return request without auth if token retrieval fails
		return c.client.R().SetContext(ctx)
	}

	return c.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+token)
}

// Close closes the client and releases resources
func (c *Client) Close() error {
	return c.client.Close()
}
