package tcnsdk

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Example demonstrates how to use the SAAS API client
func Example() {
	// Create client configuration
	config := &Config{
		BaseURL: "https://openapi1.ourvend.com", // Test environment
		AppID:   "",
		Key:     "",
		Secret:  "",
		Timeout: 30 * time.Second,
	}

	// Create client
	client := NewClient(config)
	defer client.Close()

	ctx := context.Background()
	machineID := "2412090001"

	// Example 1: Get machine information
	fmt.Println("=== Getting Machine Information ===")
	machineInfo, err := client.Machine.GetMachineInfo(ctx, machineID)
	if err != nil {
		log.Printf("Failed to get machine info: %v", err)
	} else {
		fmt.Printf("Machine: %s, Signal: %d, Temperature: %s°C, Running State: %d\n",
			machineInfo.Data.MachineName,
			machineInfo.Data.Signal,
			machineInfo.Data.Temperature,
			machineInfo.Data.RunningState)
	}

	// Example 2: Get machine commodities
	fmt.Println("\n=== Getting Machine Commodities ===")
	commodities, err := client.Machine.GetMachineCommodities(ctx, machineID)
	if err != nil {
		log.Printf("Failed to get machine commodities: %v", err)
	} else {
		fmt.Printf("Found %d commodities:\n", len(commodities.Data))
		for _, commodity := range commodities.Data {
			fmt.Printf("- %s (SKU: %s): $%.2f, Weight: %dg\n",
				commodity.CommodityName,
				commodity.CommoditySKU,
				commodity.Price,
				commodity.Weight)
		}
	}

	// Example 3: Get payment configuration
	fmt.Println("\n=== Getting Payment Configuration ===")
	payConfig, err := client.Machine.GetMachinePayConfig(ctx, machineID)
	if err != nil {
		log.Printf("Failed to get payment config: %v", err)
	} else {
		fmt.Printf("Currency: %s, Pre-auth Limit: %s, QR Code: %s\n",
			payConfig.Data.Currency,
			payConfig.Data.PreAuthorLimit,
			payConfig.Data.QRCodeURL)
	}

	// Example 4: Open door for purchase
	fmt.Println("\n=== Opening Door for Purchase ===")
	openDoorReq := &OpenDoorRequest{
		OrderID:         "ORDER123456789",
		MachineID:       machineID,
		DoorNo:          1,
		TranseType:      TranseTypePurchase,
		CustomerDetails: "customer123",
		NotifyURL:       "https://your-callback-url.com/webhook",
		Remark:          "Test purchase",
	}

	openDoorResp, err := client.Order.OpenDoor(ctx, openDoorReq)
	if err != nil {
		log.Printf("Failed to open door: %v", err)
	} else {
		fmt.Printf("Door opened: OrderID=%s, Action=%s, Status=%t, Message=%s\n",
			openDoorResp.OrderID,
			openDoorResp.Action,
			openDoorResp.Status,
			openDoorResp.Msg)
	}

	// Example 5: Open door for restocking (testing)
	fmt.Println("\n=== Opening Door for Restocking ===")
	restockOpenDoorReq := &RestockOpenDoorRequest{
		OrderID:         "RESTOCK123456789",
		MachineID:       machineID,
		DoorNo:          1,
		TranseType:      TranseTypeRestock,
		CustomerDetails: "restock123",
		NotifyURL:       "https://your-callback-url.com/webhook",
		Remark:          "Test restock",
	}

	restockOpenDoorResp, err := client.Order.RestockOpenDoor(ctx, restockOpenDoorReq)
	if err != nil {
		log.Printf("Failed to open restock door: %v", err)
	} else {
		fmt.Printf("Restock door opened: OrderID=%s, Action=%s, Status=%t, Message=%s\n",
			restockOpenDoorResp.OrderID,
			restockOpenDoorResp.Action,
			restockOpenDoorResp.Status,
			restockOpenDoorResp.Msg)
	}

	// Example 6: Report payment result
	fmt.Println("\n=== Reporting Payment Result ===")
	paymentReportReq := &ReportPaymentResultRequest{
		OrderID:   "ORDER123456789",
		OrderNo:   "233355444",
		PayType:   PayTypeNormal,
		PayStatus: PayStatusSuccess,
	}

	paymentReportResp, err := client.Order.ReportPaymentResult(ctx, paymentReportReq)
	if err != nil {
		log.Printf("Failed to report payment result: %v", err)
	} else {
		fmt.Printf("Payment reported: Status=%t, Message=%s\n",
			paymentReportResp.Status,
			paymentReportResp.Message)
	}

	// Example 7: Confirm replenishment
	fmt.Println("\n=== Confirming Replenishment ===")
	replenishmentReq := &ConfirmRepliRequest{
		Mid:        "MERCHANT123",
		TransID:    1234567890123,
		TrackingID: "TRACKING001",
		DoorNo:     1,
		Data: []ConfirmRepliItem{
			{
				CommodityID: "2164367863923397",
				LayerNo:     2,
				SpotCount:   1,
				StockCount:  500,
				Sno:         123,
			},
		},
	}

	replenishmentResp, err := client.Order.ConfirmReplenishment(ctx, replenishmentReq)
	if err != nil {
		log.Printf("Failed to confirm replenishment: %v", err)
	} else {
		fmt.Printf("Replenishment confirmed: Success=%t, Message=%s\n",
			replenishmentResp.Success,
			replenishmentResp.Message)
	}
}
