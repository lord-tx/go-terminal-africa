package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/lord-tx/go-terminal-africa"
)

func TestUserService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/users/USR-1":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.User]{
				Status: true,
				Data:   terminal.User{ID: "USR-1", Email: "test@example.com"},
			})
		case "/users/wallet":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.WalletInfo]{
				Status: true,
				Data:   terminal.WalletInfo{WalletID: "W-1", Balance: 10000},
			})
		case "/users/wallet-balance":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.WalletBalance]{
				Status: true,
				Data:   terminal.WalletBalance{Balance: 10000, Currency: "NGN"},
			})
		case "/users/carriers":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Carrier]{
				Status: true,
				Data:   []terminal.Carrier{{ID: "CR-1", Name: "DHL"}},
			})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("GetUser", func(t *testing.T) {
		res, err := client.Users.GetUser(ctx, "USR-1")
		if err != nil || !res.Status || res.Data.ID != "USR-1" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("WalletBalance", func(t *testing.T) {
		res, err := client.Users.WalletBalance(ctx, "USR-1")
		if err != nil || !res.Status || res.Data.Balance != 10000 {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("Carriers", func(t *testing.T) {
		res, err := client.Users.Carriers(ctx)
		if err != nil || !res.Status || len(res.Data) != 1 {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})
}
