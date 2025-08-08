# TCN SDK - SAAS API V2 Client

A clean, comprehensive Go client for the SAAS API V2 using [resty v3](https://resty.dev/).

## Features

- **Clean Architecture**: Well-structured, maintainable code following Go best practices
- **Type Safety**: Comprehensive type definitions for all API requests and responses
- **Error Handling**: Proper error handling with meaningful error messages
- **Service Separation**: Organized into logical services (Machine, Order, Recognition)
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

### Recognition Service

#### ConfirmCommodity
Queries cloud product inventory to confirm product information.

```go
func (s *RecognitionService) ConfirmCommodity(ctx context.Context, req *CommodityConfirmRequest) (*CommodityConfirmResponse, error)
```

#### RegisterProduct
Registers a new product for visual recognition.

```go
func (s *RecognitionService) RegisterProduct(ctx context.Context, req *ProductRegistrationRequest) (*ProductRegistrationResponse, error)
```

#### QueryProductReview
Queries the status of a product review.

```go
func (s *RecognitionService) QueryProductReview(ctx context.Context, req *ProductReviewQueryRequest) (*ProductReviewQueryResponse, error)
```

#### SubmitRecognition
Submits a video for visual recognition.

```go
func (s *RecognitionService) SubmitRecognition(ctx context.Context, req *RecognitionRequest) (*RecognitionResponse, error)
```

#### QueryRecognitionResult
Queries the result of a recognition task.

```go
func (s *RecognitionService) QueryRecognitionResult(ctx context.Context, req *RecognitionResultQueryRequest) (*RecognitionResultQueryResponse, error)
```

#### GetRemainingQuota
Gets the remaining recognition quota.

```go
func (s *RecognitionService) GetRemainingQuota(ctx context.Context, appID string) (*RemainingQuotaResponse, error)
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

### Recognition States
```go
const (
    RecognitionStatePending   = 0  // Pending review
    RecognitionStateReviewing = 1  // Under review
    RecognitionStateApproved  = 2  // Approved
    RecognitionStateRejected  = 3  // Rejected
)
```

### Recognition Result Status
```go
const (
    RecognitionResultSuccess = 2  // Recognition successful
)
```

### Resource Types
```go
const (
    ResourceTypeURL = 1  // Video URL
    ResourceTypeID  = 2  // Video resource ID
)
```

## Examples

### Callback Examples

The SDK provides comprehensive callback structs for handling webhook events from the SAAS API. Here are examples of how to use them:

#### Basic Callback Handling

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    
    "github.com/techpartners-asia/tcnsdk"
)

// Handle incoming webhook callbacks
func handleWebhook(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse the callback event
    var callbackEvent tcnsdk.CallbackEvent
    if err := json.NewDecoder(r.Body).Decode(&callbackEvent); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Handle different callback types
    switch callbackEvent.Action {
    case tcnsdk.CallbackActionOrderDetected:
        handleOrderDetected(callbackEvent)
    case tcnsdk.CallbackActionOpenedDoor:
        handleOpenedDoor(callbackEvent)
    case tcnsdk.CallbackActionCloseDoor:
        handleCloseDoor(callbackEvent)
    case tcnsdk.CallbackActionCancel:
        handleCancel(callbackEvent)
    case tcnsdk.CallbackActionOrderSettlement:
        handleOrderSettlement(callbackEvent)
    case tcnsdk.CallbackActionOrderAdjustment:
        handleOrderAdjustment(callbackEvent)
    case tcnsdk.CallbackActionOrderRefund:
        handleOrderRefund(callbackEvent)
    default:
        log.Printf("Unknown callback action: %s", callbackEvent.Action)
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"success"}`))
}

// Handle OrderDetected callback (AI recognition)
func handleOrderDetected(event tcnsdk.CallbackEvent) {
    var orderDetected tcnsdk.OrderDetectedCallback
    if err := json.Unmarshal([]byte(fmt.Sprintf("%v", event.Data)), &orderDetected.Data); err != nil {
        log.Printf("Failed to parse OrderDetected data: %v", err)
        return
    }
    
    // Copy common fields
    orderDetected.OrderID = event.OrderID
    orderDetected.OrderNo = event.OrderNo
    orderDetected.TranseType = tcnsdk.TranseType(event.TranseType)
    orderDetected.Action = event.Action
    orderDetected.Status = event.Status
    orderDetected.Msg = event.Msg

    log.Printf("OrderDetected: OrderID=%s, TranseType=%d, Status=%t", 
        orderDetected.OrderID, orderDetected.TranseType, orderDetected.Status)

    // Check transaction type
    switch orderDetected.TranseType {
    case tcnsdk.TranseTypePurchase:
        log.Println("Purchase transaction detected")
    case tcnsdk.TranseTypeRestock:
        log.Println("Restock transaction detected")
    }

    // Process detection results
    if orderDetected.Status {
        for _, product := range orderDetected.Data.TradeProductModels {
            log.Printf("Product: %s, Quantity: %d, Price: %.2f", 
                product.AlisName, product.BuyCount, product.Price)
        }
    }
}

// Handle OpenedDoor callback
func handleOpenedDoor(event tcnsdk.CallbackEvent) {
    openedDoor := tcnsdk.OpenedDoorCallback{
        OrderID:         event.OrderID,
        OrderNo:         event.OrderNo,
        TranseType:      tcnsdk.TranseType(event.TranseType),
        Action:          event.Action,
        Status:          event.Status,
        Msg:             event.Msg,
        CustomerDetails: event.CustomerDetails,
        Data:            fmt.Sprintf("%v", event.Data),
    }

    log.Printf("Door opened: OrderID=%s, Status=%t, CustomerDetails=%s", 
        openedDoor.OrderID, openedDoor.Status, openedDoor.CustomerDetails)
}

// Handle CloseDoor callback
func handleCloseDoor(event tcnsdk.CallbackEvent) {
    closeDoor := tcnsdk.CloseDoorCallback{
        OrderID:         event.OrderID,
        OrderNo:         event.OrderNo,
        TranseType:      tcnsdk.TranseType(event.TranseType),
        Action:          event.Action,
        Status:          event.Status,
        Msg:             event.Msg,
        CustomerDetails: event.CustomerDetails,
        Data:            fmt.Sprintf("%v", event.Data),
    }

    log.Printf("Door closed: OrderID=%s, Status=%t", 
        closeDoor.OrderID, closeDoor.Status)
}

// Handle Cancel callback
func handleCancel(event tcnsdk.CallbackEvent) {
    cancel := tcnsdk.CancelCallback{
        OrderID:         event.OrderID,
        OrderNo:         event.OrderNo,
        TranseType:      tcnsdk.TranseType(event.TranseType),
        Action:          event.Action,
        Status:          event.Status,
        Msg:             event.Msg,
        CustomerDetails: event.CustomerDetails,
        Data:            fmt.Sprintf("%v", event.Data),
    }

    log.Printf("Order cancelled: OrderID=%s, Reason=%s", 
        cancel.OrderID, cancel.Msg)
}

// Handle OrderSettlement callback
func handleOrderSettlement(event tcnsdk.CallbackEvent) {
    var settlement tcnsdk.OrderSettlementCallback
    if err := json.Unmarshal([]byte(fmt.Sprintf("%v", event.Data)), &settlement.Data); err != nil {
        log.Printf("Failed to parse OrderSettlement data: %v", err)
        return
    }
    
    settlement.OrderID = event.OrderID
    settlement.OrderNo = event.OrderNo
    settlement.TranseType = tcnsdk.TranseType(event.TranseType)
    settlement.Action = event.Action
    settlement.Status = event.Status
    settlement.Msg = event.Msg
    settlement.CustomerDetails = event.CustomerDetails

    log.Printf("Order settled: OrderID=%s, ProcessMode=%d", 
        settlement.OrderID, settlement.Data.TradeProcessMode)
}

// Handle OrderAdjustment callback
func handleOrderAdjustment(event tcnsdk.CallbackEvent) {
    var adjustment tcnsdk.OrderAdjustmentCallback
    if err := json.Unmarshal([]byte(fmt.Sprintf("%v", event.Data)), &adjustment.Data); err != nil {
        log.Printf("Failed to parse OrderAdjustment data: %v", err)
        return
    }
    
    adjustment.OrderID = event.OrderID
    adjustment.OrderNo = event.OrderNo
    adjustment.TranseType = tcnsdk.TranseType(event.TranseType)
    adjustment.Action = event.Action
    adjustment.Status = event.Status
    adjustment.Msg = event.Msg
    adjustment.CustomerDetails = event.CustomerDetails

    log.Printf("Order adjusted: OrderID=%s, AdjustmentOrderNo=%s", 
        adjustment.OrderID, adjustment.Data.OrderNo)
}

// Handle OrderRefund callback
func handleOrderRefund(event tcnsdk.CallbackEvent) {
    var refund tcnsdk.OrderRefundCallback
    if err := json.Unmarshal([]byte(fmt.Sprintf("%v", event.Data)), &refund.Data); err != nil {
        log.Printf("Failed to parse OrderRefund data: %v", err)
        return
    }
    
    refund.OrderID = event.OrderID
    refund.OrderNo = event.OrderNo
    refund.TranseType = tcnsdk.TranseType(event.TranseType)
    refund.Action = event.Action
    refund.Status = event.Status
    refund.Msg = event.Msg
    refund.CustomerDetails = event.CustomerDetails

    log.Printf("Order refunded: OrderID=%s, RefundOrderNo=%s, Reason=%s", 
        refund.OrderID, refund.Data.RefundOrderNo, refund.Data.RefundRemark)
}

func main() {
    // Set up webhook server
    http.HandleFunc("/webhook", handleWebhook)
    
    log.Println("Starting webhook server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

#### Visual Recognition Callback Examples

```go
// Handle ProductReviewNotification callback
func handleProductReview(callback tcnsdk.ProductReviewNotificationCallback) {
    log.Printf("Product review: YsSkuId=%s, State=%d", 
        callback.YsSkuId, callback.State)

    switch callback.State {
    case tcnsdk.RecognitionStatePending:
        log.Println("Product review pending")
    case tcnsdk.RecognitionStateReviewing:
        log.Println("Product under review")
    case tcnsdk.RecognitionStateApproved:
        log.Println("Product approved")
    case tcnsdk.RecognitionStateRejected:
        log.Printf("Product rejected: %s", callback.Desc)
    }
}

// Handle RecognitionCallback
func handleRecognitionResult(callback tcnsdk.RecognitionCallback) {
    log.Printf("Recognition result: TaskId=%s, State=%d", 
        callback.TaskId, callback.State)

    if callback.ResultStatus.Code == tcnsdk.RecognitionResultSuccess {
        log.Println("Recognition successful")
        for _, item := range callback.ResultData {
            log.Printf("Product: %s, Quantity: %d", 
                item.CommoditySku, item.Qty)
        }
    } else {
        log.Printf("Recognition failed: %s", callback.ResultStatus.Desc)
    }
}
```

#### Using Type-Safe Constants

```go
// Opening door with proper types
openDoorReq := &tcnsdk.OpenDoorRequest{
    OrderID:         "ORDER123456789",
    MachineID:       "2412090001",
    DoorNo:          1,
    TranseType:      tcnsdk.TranseTypePurchase,  // Type-safe constant
    CustomerDetails: "customer123",
    NotifyURL:       "https://your-callback-url.com/webhook",
    Remark:          "Test purchase",
}

// Reporting payment result with proper types
paymentReportReq := &tcnsdk.ReportPaymentResultRequest{
    OrderID:   "ORDER123456789",
    OrderNo:   "233355444",
    PayType:   tcnsdk.PayTypeNormal,     // Type-safe constant
    PayStatus: tcnsdk.PayStatusSuccess,  // Type-safe constant
}

// Submitting recognition with proper types
recognitionReq := &tcnsdk.RecognitionRequest{
    AppID:        "your-app-id",
    TaskId:       "TASK123456",
    ResourceType: tcnsdk.ResourceTypeURL,  // Type-safe constant
    ResourceUrl:  []string{"https://example.com/video.mp4"},
    ProductRange: []string{"YSK123456789"},
    NotifyUrl:    "https://your-callback-url.com/recognition",
}
```

### Complete Example

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

    // Open door for purchase
    openDoorReq := &tcnsdk.OpenDoorRequest{
        OrderID:         "ORDER123456789",
        MachineID:       machineID,
        DoorNo:          1,
        TranseType:      tcnsdk.TranseTypePurchase,
        CustomerDetails: "customer123",
        NotifyURL:       "https://your-callback-url.com/webhook",
        Remark:          "Test purchase",
    }

    openDoorResp, err := client.Order.OpenDoor(ctx, openDoorReq)
    if err != nil {
        log.Printf("Failed to open door: %v", err)
    } else {
        log.Printf("Door opened: OrderID=%s, Action=%s, Status=%t", 
            openDoorResp.OrderID, openDoorResp.Action, openDoorResp.Status)
    }

    // Report payment result
    paymentReportReq := &tcnsdk.ReportPaymentResultRequest{
        OrderID:   "ORDER123456789",
        OrderNo:   "233355444",
        PayType:   tcnsdk.PayTypeNormal,
        PayStatus: tcnsdk.PayStatusSuccess,
    }

    paymentReportResp, err := client.Order.ReportPaymentResult(ctx, paymentReportReq)
    if err != nil {
        log.Printf("Failed to report payment result: %v", err)
    } else {
        log.Printf("Payment reported: Status=%t, Message=%s", 
            paymentReportResp.Status, paymentReportResp.Message)
    }
}
```

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