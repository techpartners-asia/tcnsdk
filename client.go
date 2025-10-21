package tcnsdk

import (
	"time"

	"resty.dev/v3"

	"github.com/techpartners-asia/tcnsdk/structs"
)

// Client represents the SAAS API client
type Client struct {
	client *resty.Client
	config *Config

	// Services
	Machine *MachineService
	Order   *OrderService
	Product *ProductService
	Train   *TrainService
}

// Config holds the client configuration
type Config struct {
	BaseURL string
	AppID   string
	Key     string
	Secret  string
	Timeout time.Duration
	Debug   bool
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseURL: "https://openapi1.ourvend.com",
		Timeout: 30 * time.Second,
		Debug:   false,
	}
}

// NewClient creates a new SAAS API client
func NewClient(config *Config) *Client {
	if config == nil {
		config = DefaultConfig()
	}

	client := resty.New().
		SetBaseURL(config.BaseURL).
		SetTimeout(config.Timeout).
		SetDebug(config.Debug)

	c := &Client{
		client: client,
		config: config,
	}

	// Initialize services
	c.Machine = &MachineService{Client: c}
	c.Order = &OrderService{Client: c}
	c.Product = &ProductService{client: c}
	c.Train = &TrainService{client: c}

	return c
}

// getAuthToken retrieves authentication token
func (c *Client) getAuthResponse() (*structs.AuthResponse, error) {
	authReq := &structs.AuthRequest{
		AppID:  c.config.AppID,
		Key:    c.config.Key,
		Secret: c.config.Secret,
	}

	var authResp structs.AuthResponse
	_, err := c.client.R().
		SetBody(authReq).
		SetResult(&authResp).
		Post("/OpenApi/Login")

	if err != nil {
		return nil, err
	}

	return &authResp, nil
}

// Request performs an authenticated request
func (c *Client) Request() *resty.Request {
	authResp, err := c.getAuthResponse()
	if err != nil {
		// Return request without auth if token retrieval fails
		return c.client.R()
	}

	return c.client.R().
		SetHeader("Authorization", authResp.Data.Token)
}

// Close closes the client and releases resources
func (c *Client) Close() error {
	return c.client.Close()
}
