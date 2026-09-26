package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/danielozeh/go-terminal-africa"
)

func TestCarrierService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/carriers":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Carrier]{
				Status: true,
				Data:   []terminal.Carrier{{ID: "CR-1", Name: "DHL"}},
			})
		case "/carriers/CR-1":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Carrier]{
				Status: true,
				Data:   terminal.Carrier{ID: "CR-1", Name: "DHL"},
			})
		case "/carriers/enable/CR-1":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Carrier]{
				Status: true,
				Data:   terminal.Carrier{ID: "CR-1", Active: true},
			})
		case "/carriers/multiple/enable":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Carrier]{
				Status: true,
				Data:   []terminal.Carrier{{ID: "CR-1", Active: true}},
			})
		default:
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Carrier]{Status: true})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("GetCarriers", func(t *testing.T) {
		active := true
		res, err := client.Carriers.GetCarriers(ctx, &terminal.CarrierListParams{Active: &active})
		if err != nil || !res.Status || len(res.Data) != 1 {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("GetCarrier", func(t *testing.T) {
		res, err := client.Carriers.GetCarrier(ctx, "CR-1")
		if err != nil || !res.Status || res.Data.Name != "DHL" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("EnableCarrier", func(t *testing.T) {
		res, err := client.Carriers.EnableCarrier(ctx, "CR-1", &terminal.CarrierToggleParams{Domestic: true})
		if err != nil || !res.Status {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("EnableMultipleCarriers", func(t *testing.T) {
		res, err := client.Carriers.EnableMultipleCarriers(ctx, &terminal.EnableMultipleCarriersParams{
			Carriers: []terminal.CarrierItemToggle{{CarrierID: "CR-1", Domestic: true}},
		})
		if err != nil || !res.Status {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})
}
