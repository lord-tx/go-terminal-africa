package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/danielozeh/go-terminal-africa"
)

func TestRateService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/rates/shipment":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Rate]{
				Status: true,
				Data:   []terminal.Rate{{ID: "RT-1", Amount: 1500, CarrierName: "DHL"}},
			})
		case "/rates/shipment/quotes":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Rate]{
				Status: true,
				Data:   []terminal.Rate{{ID: "RT-QUOTE", Amount: 2000}},
			})
		case "/rates/RT-1":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Rate]{
				Status: true,
				Data:   terminal.Rate{ID: "RT-1", Amount: 1500},
			})
		default:
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Rate]{Status: true})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("GetShipmentRates", func(t *testing.T) {
		res, err := client.Rates.GetShipmentRates(ctx, &terminal.ShipmentRateParams{
			PickupAddress: "AD-1", DeliveryAddress: "AD-2", ParcelID: "PC-1",
		})
		if err != nil || !res.Status || len(res.Data) != 1 {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("GetQuotesForShipment", func(t *testing.T) {
		res, err := client.Rates.GetQuotesForShipment(ctx, &terminal.QuotesParams{
			PickupAddress: map[string]any{"city": "Lagos"}, DeliveryAddress: map[string]any{"city": "Ikeja"},
		})
		if err != nil || !res.Status || len(res.Data) != 1 {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("GetRate", func(t *testing.T) {
		res, err := client.Rates.GetRate(ctx, "RT-1")
		if err != nil || !res.Status || res.Data.ID != "RT-1" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})
}
