package tcnsdk

import (
	"context"
	"net/http"
	"time"

	"resty.dev/v3"

	"github.com/techpartners-asia/tcnsdk/structs"
)

const (
	defaultBaseURL = "https://openapi1.ourvend.com"
	defaultTimeout = 30 * time.Second
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
	BaseURL    string
	AppID      string
	Key        string
	Secret     string
	Timeout    time.Duration
	Debug      bool
	HTTPClient *http.Client
}

// Option defines a functional option for configuring the client
type Option func(*Config)

// WithBaseURL sets the API base URL
func WithBaseURL(url string) Option {
	return func(c *Config) {
		c.BaseURL = url
	}
}

// WithTimeout sets the request timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithDebug enables debug mode
func WithDebug(debug bool) Option {
	return func(c *Config) {
		c.Debug = debug
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(client *http.Client) Option {
	return func(c *Config) {
		c.HTTPClient = client
	}
}

// NewClient creates a new SAAS API client with the given credentials and options
func NewClient(appID, key, secret string, opts ...Option) *Client {
	config := &Config{
		BaseURL: defaultBaseURL,
		AppID:   appID,
		Key:     key,
		Secret:  secret,
		Timeout: defaultTimeout,
	}

	for _, opt := range opts {
		opt(config)
	}

	var client *resty.Client
	if config.HTTPClient != nil {
		client = resty.NewWithClient(config.HTTPClient).
			SetBaseURL(config.BaseURL).
			SetTimeout(config.Timeout).
			SetDebug(config.Debug)
	} else {
		client = resty.New().
			SetBaseURL(config.BaseURL).
			SetTimeout(config.Timeout).
			SetDebug(config.Debug)
	}

	c := &Client{
		client: client,
		config: config,
	}

	// Initialize services
	c.Machine = &MachineService{client: c}
	c.Order = &OrderService{client: c}
	c.Product = &ProductService{client: c}
	c.Train = &TrainService{client: c}

	return c
}

// getAuthResponse retrieves authentication token
func (c *Client) getAuthResponse(ctx context.Context) (*structs.AuthResponse, error) {
	authReq := &structs.AuthRequest{
		AppID:  c.config.AppID,
		Key:    c.config.Key,
		Secret: c.config.Secret,
	}

	var authResp structs.AuthResponse
	_, err := c.client.R().
		SetContext(ctx).
		SetBody(authReq).
		SetResult(&authResp).
		Post("/OpenApi/Login")

	if err != nil {
		return nil, err
	}

	return &authResp, nil
}

// Close closes the client and releases resources
func (c *Client) Close() error {
	// Resty client doesn't have a Close method, but we keep this for future extensibility
	// or if we need to close underlying connections explicitly
	return nil
}
