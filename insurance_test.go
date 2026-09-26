package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/lord-tx/go-terminal-africa"
)

func TestInsuranceService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/insurance":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Insurance]{
				Status: true,
				Data:   []terminal.Insurance{{ID: "INS-1", Amount: 500}},
			})
		case "/insurance/premium":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.PremiumResponse]{
				Status: true,
				Data:   terminal.PremiumResponse{Premium: 100, Value: 5000},
			})
		case "/insurance/purchase":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Insurance]{
				Status: true,
				Data:   terminal.Insurance{ID: "INS-1", Status: "active"},
			})
		default:
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Insurance]{
				Status: true,
				Data:   terminal.Insurance{ID: "INS-1"},
			})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("GetInsuranceList", func(t *testing.T) {
		res, err := client.Insurance.GetInsuranceList(ctx, &terminal.InsuranceListParams{Page: 1})
		if err != nil || !res.Status || len(res.Data) != 1 {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("GetInsurancePremium", func(t *testing.T) {
		res, err := client.Insurance.GetInsurancePremium(ctx, "PC-100")
		if err != nil || !res.Status || res.Data.Premium != 100 {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("PurchaseInsurance", func(t *testing.T) {
		res, err := client.Insurance.PurchaseInsurance(ctx, "SH-100")
		if err != nil || !res.Status || res.Data.ID != "INS-1" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})
}
