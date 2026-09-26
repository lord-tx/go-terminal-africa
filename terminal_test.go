package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	terminal "github.com/lord-tx/go-terminal-africa"
)

func TestNewClient(t *testing.T) {
	client := terminal.NewClient("sk_test_123", terminal.WithTimeout(10*time.Second))
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestDefaultClientAndInit(t *testing.T) {
	client := terminal.Init("sk_test_global")
	if client == nil {
		t.Fatal("expected non-nil initialized client")
	}
}

func TestAPIErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"status":  false,
			"message": "Invalid address payload",
		})
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	_, err := client.Addresses.GetAddress(context.Background(), "invalid_id")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
