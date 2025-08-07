# TCN SDK - SAAS API V2 Client

A clean, comprehensive Go client for the SAAS API V2 using [resty v3](https://resty.dev/).

## Features

- **Clean Architecture**: Well-structured, maintainable code following Go best practices
- **Type Safety**: Comprehensive type definitions for all API requests and responses
- **Error Handling**: Proper error handling with meaningful error messages
- **Service Separation**: Organized into logical services (Machine, Order)
- **Context Support**: Full context support for cancellation and timeouts
- **Constants**: Named constants for all enum values

## Installation

```bash
go get github.com/techpartners-asia/tcnsdk
```

## Quick Start

```go
package main

import (
    "context"
    "log"
    "time"
    
    "github.com/techpartners-asia/tcnsdk"
)

func main() {
    // Create client configuration
    config := &tcnsdk.Config{
        BaseURL: "https://openapi1.ourvend.com", // Test environment
        AppID:   "your-app-id",
        Key:     "your-key",
        Secret:  "your-secret",
        Timeout: 30 * time.Second,
    }

    // Create client
    client := tcnsdk.NewClient(config)
    defer client.Close()

    ctx := context.Background()
    machineID := "2412090001"

    // Get machine information
    machineInfo, err := client.Machine.GetMachineInfo(ctx, machineID)
    if err != nil {
        log.Fatalf("Failed to get machine info: %v", err)
    }

    log.Printf("Machine: %s, Signal: %d", 
        machineInfo.Data.MachineName, 
        machineInfo.Data.Signal)
}
```

## API Reference

### Client Configuration

```go
type Config struct {
    BaseURL string        // API base URL
    AppID   string        // Application ID
    Key     string        // API Key
    Secret  string        // API Secret
    Timeout time.Duration // Request timeout
}
```

### Machine Service

#### GetMachineInfo
Retrieves device information.

```go
func (s *MachineService) GetMachineInfo(ctx context.Context, machineID string) (*MachineInfoResponse, error)
```

#### GetMachineCommodities
Retrieves equipment product information.

```go
func (s *MachineService) GetMachineCommodities(ctx context.Context, machineID string) (*CommodityResponse, error)
```

#### GetMachinePayConfig
Retrieves device payment configuration information.

```go
func (s *MachineService) GetMachinePayConfig(ctx context.Context, machineID string) (*PayConfigResponse, error)
```

### Order Service

#### OpenDoor
Opens the vending machine door for purchase.

```go
func (s *OrderService) OpenDoor(ctx context.Context, req *OpenDoorRequest) (*OpenDoorResponse, error)
```

#### RestockOpenDoor
Opens the vending machine door for restocking (testing endpoint).

```go
func (s *OrderService) RestockOpenDoor(ctx context.Context, req *RestockOpenDoorRequest) (*RestockOpenDoorResponse, error)
```

#### ReportPaymentResult
Reports the payment result to the API.

```go
func (s *OrderService) ReportPaymentResult(ctx context.Context, req *ReportPaymentResultRequest) (*ReportPaymentResultResponse, error)
```

#### ConfirmReplenishment
Confirms replenishment data.

```go
func (s *OrderService) ConfirmReplenishment(ctx context.Context, req *ConfirmRepliRequest) (*ConfirmRepliResponse, error)
```

## Constants

### Transaction Types
```go
const (
    TranseTypePurchase = 0  // Purchase transaction
    TranseTypeRestock  = 2  // Restock transaction
)
```

### Payment Types
```go
const (
    PayTypeNormal    = 0  // Normal payment
    PayTypeDeduction = 1  // Deduction payment
    PayTypeRefund    = 2  // Refund payment
)
```

### Payment Status
```go
const (
    PayStatusSuccess = 1  // Payment successful
    PayStatusFailed  = 2  // Payment failed
)
```

### Trade Process Modes
```go
const (
    TradeProcessModeNormal    = 0  // Normal transaction
    TradeProcessModeCancel    = 1  // Cancel transaction
    TradeProcessModeInterrupt = 2  // Interrupt transaction
)
```

### Meter Types
```go
const (
    MeterTypePiece  = 1  // Piece-based pricing
    MeterTypeWeight = 2  // Weight-based pricing
)
```

## Examples

### Complete Example

See `example.go` for a complete example demonstrating all API operations.

### Authentication

The client automatically handles authentication by retrieving a token for each request:

```go
// Authentication is handled automatically
client := tcnsdk.NewClient(config)
machineInfo, err := client.Machine.GetMachineInfo(ctx, machineID)
```

### Error Handling

```go
machineInfo, err := client.Machine.GetMachineInfo(ctx, machineID)
if err != nil {
    log.Printf("Failed to get machine info: %v", err)
    return
}
```

### Context Usage

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

machineInfo, err := client.Machine.GetMachineInfo(ctx, machineID)
```

## Environment URLs

- **Test Environment**: `https://openapi1.ourvend.com`
- **Production Environment**: `https://openapi.aivendortech.com`

## Dependencies

- [resty v3](https://resty.dev/) - HTTP client library
- Go 1.24.4+

## License

This project is licensed under the MIT License. 