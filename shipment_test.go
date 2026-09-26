package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/lord-tx/go-terminal-africa"
)

func TestShipmentService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/shipments":
			if r.Method == http.MethodPost {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Shipment]{
					Status: true,
					Data:   terminal.Shipment{ID: "SH-100", Status: "created"},
				})
			} else if r.Method == http.MethodGet {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Shipment]{
					Status: true,
					Data:   []terminal.Shipment{{ID: "SH-100"}},
				})
			} else if r.Method == http.MethodDelete {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[map[string]any]{
					Status: true,
					Data:   map[string]any{"deleted": true},
				})
			}
		case "/shipments/quick":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Shipment]{
				Status: true,
				Data:   terminal.Shipment{ID: "SH-QUICK"},
			})
		case "/shipments/track/SH-100":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.TrackingInfo]{
				Status: true,
				Data:   terminal.TrackingInfo{ShipmentID: "SH-100", Status: "in_transit"},
			})
		case "/shipments/cancel":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Shipment]{
				Status: true,
				Data:   terminal.Shipment{ID: "SH-100", Status: "cancelled"},
			})
		default:
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Shipment]{Status: true, Data: terminal.Shipment{ID: "SH-100"}})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("CreateShipment", func(t *testing.T) {
		res, err := client.Shipments.CreateShipment(ctx, &terminal.CreateShipmentParams{
			AddressFrom: "AD-1", AddressTo: "AD-2", Parcel: "PC-1", ShipmentPurpose: "personal",
		})
		if err != nil || !res.Status || res.Data.ID != "SH-100" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("CreateQuickShipment", func(t *testing.T) {
		res, err := client.Shipments.CreateQuickShipment(ctx, &terminal.QuickShipmentParams{
			PickupAddress: map[string]any{"city": "Lagos"}, DeliveryAddress: map[string]any{"city": "Ikeja"},
		})
		if err != nil || !res.Status || res.Data.ID != "SH-QUICK" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("TrackShipment", func(t *testing.T) {
		res, err := client.Shipments.TrackShipment(ctx, "SH-100")
		if err != nil || !res.Status || res.Data.Status != "in_transit" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("CancelShipment", func(t *testing.T) {
		res, err := client.Shipments.CancelShipment(ctx, "SH-100")
		if err != nil || !res.Status || res.Data.Status != "cancelled" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("DeleteShipment", func(t *testing.T) {
		res, err := client.Shipments.DeleteShipment(ctx, "SH-100")
		if err != nil || !res.Status {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})
}
