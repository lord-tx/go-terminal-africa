package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/lord-tx/go-terminal-africa"
)

func TestWebhookService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/webhooks":
			if r.Method == http.MethodPost {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Webhook]{
					Status: true,
					Data:   terminal.Webhook{ID: "WH-1", URL: "https://example.com/webhook"},
				})
			} else {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Webhook]{
					Status: true,
					Data:   []terminal.Webhook{{ID: "WH-1"}},
				})
			}
		case "/webhooks/WH-1":
			if r.Method == http.MethodDelete {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[map[string]any]{
					Status: true,
					Data:   map[string]any{"deleted": true},
				})
			} else {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Webhook]{
					Status: true,
					Data:   terminal.Webhook{ID: "WH-1"},
				})
			}
		case "/webhooks/enable/WH-1":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Webhook]{
				Status: true,
				Data:   terminal.Webhook{ID: "WH-1", Active: true},
			})
		case "/webhooks/disable/WH-1":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Webhook]{
				Status: true,
				Data:   terminal.Webhook{ID: "WH-1", Active: false},
			})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("CreateWebhook", func(t *testing.T) {
		res, err := client.Webhooks.CreateWebhook(ctx, &terminal.CreateWebhookParams{
			Name: "Production Webhook", URL: "https://example.com/webhook", Active: true, Events: []string{"shipment.created"}, Live: true,
		})
		if err != nil || !res.Status || res.Data.ID != "WH-1" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("EnableWebhook", func(t *testing.T) {
		res, err := client.Webhooks.EnableWebhook(ctx, "WH-1")
		if err != nil || !res.Status || !res.Data.Active {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("DisableWebhook", func(t *testing.T) {
		res, err := client.Webhooks.DisableWebhook(ctx, "WH-1")
		if err != nil || !res.Status || res.Data.Active {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("DeleteWebhook", func(t *testing.T) {
		res, err := client.Webhooks.DeleteWebhook(ctx, "WH-1")
		if err != nil || !res.Status {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})
}
