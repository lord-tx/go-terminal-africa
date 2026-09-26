package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/danielozeh/go-terminal-africa"
)

func TestParcelService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/parcels":
			if r.Method == http.MethodPost {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Parcel]{
					Status: true,
					Data:   terminal.Parcel{ID: "PC-100", Description: "Electronics"},
				})
			} else {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Parcel]{
					Status: true,
					Data:   []terminal.Parcel{{ID: "PC-100"}},
				})
			}
		case "/parcels/PC-100":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Parcel]{
				Status: true,
				Data:   terminal.Parcel{ID: "PC-100", Description: "Electronics"},
			})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("CreateParcel", func(t *testing.T) {
		res, err := client.Parcels.CreateParcel(ctx, &terminal.CreateParcelParams{
			Description: "Electronics", WeightUnit: "kg", Items: []terminal.ParcelItem{{Name: "Laptop", Weight: 2}},
		})
		if err != nil || !res.Status || res.Data.ID != "PC-100" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("GetParcel", func(t *testing.T) {
		res, err := client.Parcels.GetParcel(ctx, "PC-100")
		if err != nil || !res.Status || res.Data.ID != "PC-100" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})
}
