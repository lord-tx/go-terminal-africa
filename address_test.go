package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/danielozeh/go-terminal-africa"
)

func TestAddressService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/addresses":
			if r.Method == http.MethodPost {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Address]{
					Status: true,
					Data:   terminal.Address{ID: "AD-12345", City: "Ikeja", State: "Lagos"},
				})
			} else if r.Method == http.MethodGet {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Address]{
					Status: true,
					Data:   []terminal.Address{{ID: "AD-12345"}},
				})
			}
		case "/addresses/AD-12345":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Address]{
				Status: true,
				Data:   terminal.Address{ID: "AD-12345"},
			})
		case "/addresses/validate":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.AddressValidationResult]{
				Status: true,
				Data:   terminal.AddressValidationResult{IsValid: true},
			})
		case "/addresses/default/sender":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Address]{
				Status: true,
				Data:   terminal.Address{ID: "AD-DEFAULT"},
			})
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("CreateAddress", func(t *testing.T) {
		res, err := client.Addresses.CreateAddress(ctx, &terminal.CreateAddressParams{
			Line1: "123 Street", City: "Ikeja", State: "Lagos", Country: "NG",
		})
		if err != nil || !res.Status || res.Data.ID != "AD-12345" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("GetAddresses", func(t *testing.T) {
		res, err := client.Addresses.GetAddresses(ctx, &terminal.ListParams{Page: 1, PerPage: 10})
		if err != nil || !res.Status || len(res.Data) != 1 {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("GetAddress", func(t *testing.T) {
		res, err := client.Addresses.GetAddress(ctx, "AD-12345")
		if err != nil || !res.Status || res.Data.ID != "AD-12345" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("ValidateAddress", func(t *testing.T) {
		res, err := client.Addresses.ValidateAddress(ctx, &terminal.ValidateAddressParams{
			Country: "NG", State: "Lagos", City: "Ikeja",
		})
		if err != nil || !res.Status || !res.Data.IsValid {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("SetDefaultSenderAddress", func(t *testing.T) {
		res, err := client.Addresses.SetDefaultSenderAddress(ctx, "AD-12345")
		if err != nil || !res.Status {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("GetDefaultSenderAddress", func(t *testing.T) {
		res, err := client.Addresses.GetDefaultSenderAddress(ctx)
		if err != nil || !res.Status || res.Data.ID != "AD-DEFAULT" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})
}
