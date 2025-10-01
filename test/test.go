package main

import (
	"context"
	"fmt"
	"time"

	"github.com/techpartners-asia/tcnsdk"
	"github.com/techpartners-asia/tcnsdk/structs"
)

func main() {
	// Your credentials
	config := &tcnsdk.Config{
		BaseURL: "https://openapi1.ourvend.com",
		AppID:   "P537217083646021",
		Key:     "loYgqEZkCUuDItrnxHvn",
		Secret:  "fmYGOcX4C3InYk3wqRaZb2Z9jWlKaU",
		Timeout: 30 * time.Second,
		Debug:   true, // Set to false for cleaner output
	}

	client := tcnsdk.NewClient(config)
	defer client.Close()

	ctx := context.Background()
	machineID := "2504150004"

	fmt.Printf("🚀 Testing TCN SDK\n")
	fmt.Printf("Machine: %s\n", machineID)
	fmt.Println("==================================================")

	// Test machine info
	fmt.Println("\n📱 Getting machine info...")
	info, err := client.Machine.GetMachineInfo(ctx, machineID)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Machine: %s\n", info.Data.MachineName)
	}

	// Test products
	fmt.Println("\n📦 Getting products...")
	products, err := client.Product.ListProducts(ctx, &structs.ProductListRequest{
		PageIndex: 1,
		PageSize:  5,
	})
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Found %d products\n", len(products.Data.Items))
	}

	fmt.Println("\n🎉 Test completed!")
}
