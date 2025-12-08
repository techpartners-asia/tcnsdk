# TCN SDK for Go

`tcnsdk` is a Go client library for interacting with the SAAS Transaction Related Interface V2.1.2. It provides a structured and idiomatic way to access endpoints for machine management, order processing, product catalog, and more.

## Features

- **Complete API Coverage**: Support for Authentication, Product Management, Equipment Management, Operation & Sales, and Callbacks.
- **Context Support**: All API methods accept `context.Context` for timeout and cancellation control.
- **Type-Safe Structs**: Request and response payloads are mapped to Go structs for easy usage.
- **Functional Options**: Configure the client with custom HTTP clients, timeouts, and debug mode.

## Installation

```bash
go get github.com/techpartners-asia/tcnsdk
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/techpartners-asia/tcnsdk"
	"github.com/techpartners-asia/tcnsdk/structs"
)

func main() {
	// 1. Initialize the client
	client := tcnsdk.NewClient(
		"your-app-id",
		"your-key",
		"your-secret",
		tcnsdk.WithTimeout(10*time.Second),
		tcnsdk.WithDebug(true),
	)

	ctx := context.Background()

	// 2. List Machines
	machines, err := client.Machine.ListMachines(ctx, &structs.MachineListRequest{
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		log.Fatalf("Failed to list machines: %v", err)
	}

	fmt.Printf("Found %d machines\n", machines.Data.Total)
	for _, m := range machines.Data.Items {
		fmt.Printf("- %s (%s)\n", m.MachineName, m.MachineID)
	}

	// 3. Open a Door
	resp, err := client.Order.OpenDoor(ctx, &structs.OpenDoorRequest{
		OrderID:    "order-12345",
		MachineID:  "machine-001",
		DoorNo:     1,
		TranseType: structs.TranseTypePurchase,
		NotifyURL:  "https://your-callback.com/notify",
	})
	if err != nil {
		log.Fatalf("Failed to open door: %v", err)
	}
	fmt.Printf("Open Door Status: %v\n", resp.Succeeded)
}
```

## Configuration

You can customize the client using functional options:

- `WithBaseURL(url string)`: Override the default API endpoint.
- `WithTimeout(d time.Duration)`: Set global request timeout.
- `WithDebug(enable bool)`: Enable HTTP request/response logging.
- `WithHTTPClient(client *http.Client)`: Bring your own `http.Client`.

## Documentation

For detailed API documentation, please refer to the official [SAAS Transaction Related Interface Documentation](https://alidocs.dingtalk.com/i/p/K1nRzl9ogbvXbxZq7kmbLW4EEBB3gzLq).

## License

[License Name]

