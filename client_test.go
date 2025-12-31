package tcnsdk

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/techpartners-asia/tcnsdk/structs"
)

func TestClient_GetAuthResponse(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/OpenApi/Login" {
			t.Errorf("Expected path /OpenApi/Login, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"statusCode": 200,
			"data": {
				"token": "mock-token",
				"expiresIn": 3600
			},
			"succeeded": true
		}`))
	}))
	defer server.Close()

	// Initialize client with mock server URL
	client := NewClient("appID", "key", "secret", WithBaseURL(server.URL))

	// Call the method
	resp, err := client.getAuthResponse(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Assertions
	if resp.Data.Token != "mock-token" {
		t.Errorf("Expected token 'mock-token', got %s", resp.Data.Token)
	}
}

func TestMachineService_ListMachines(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock Login first if needed, or assume token is already set if we manually set it (Resty client state)
		// But here getAuthResponse is called internally. Let's mock both endpoints.
		if r.URL.Path == "/OpenApi/Login" {
			w.Write([]byte(`{"statusCode": 200, "data": {"token": "mock"}, "succeeded": true}`))
			return
		}
		if r.URL.Path == "/OpenApi/Machines" {
			w.Write([]byte(`{
				"statusCode": 200,
				"data": {
					"total": 1,
					"items": [{"machineId": "M001", "machineName": "Test Machine"}]
				},
				"succeeded": true
			}`))
			return
		}
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer server.Close()

	client := NewClient("id", "key", "sec", WithBaseURL(server.URL))

	req := &structs.MachineListRequest{PageIndex: 1, PageSize: 10}
	resp, err := client.Machine.ListMachines(context.Background(), req)
	if err != nil {
		t.Fatalf("ListMachines failed: %v", err)
	}

	if resp.Data.Total != 1 {
		t.Errorf("Expected 1 machine, got %d", resp.Data.Total)
	}
	if resp.Data.Items[0].MachineID != "M001" {
		t.Errorf("Expected machine ID M001, got %s", resp.Data.Items[0].MachineID)
	}
}

func TestMachineService_AddProductToMachine(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/OpenApi/Login" {
			w.Write([]byte(`{"statusCode": 200, "data": {"token": "mock"}, "succeeded": true}`))
			return
		}
		if r.URL.Path == "/OpenApi/VendSlotCommoditys/Add" {
			// Verify body is an array
			body := make([]byte, r.ContentLength)
			r.Body.Read(body)

			// Simple check if it starts with [
			if len(body) > 0 && body[0] != '[' {
				t.Errorf("Expected body to start with [, got %s", string(body))
			}

			w.Write([]byte(`{
				"statusCode": 200,
				"succeeded": true
			}`))
			return
		}
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer server.Close()

	client := NewClient("id", "key", "sec", WithBaseURL(server.URL))

	req := &structs.AddProductToMachineRequest{
		{
			VendID:          "2504150004",
			CommodityID:     "2859854396900037",
			LayerNo:         2,
			DoorNo:          1,
			Capacity:        20,
			EarlyRate:       true,
			EarlyWarigCount: 5,
			Price:           2,
		},
	}
	_, err := client.Machine.AddProductToMachine(context.Background(), req)
	if err != nil {
		t.Fatalf("AddProductToMachine failed: %v", err)
	}
}
